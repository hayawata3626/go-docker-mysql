package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// User ...
type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	err := godotenv.Load()
	r := gin.Default()
	db := connectDB()
	db.AutoMigrate(&User{})
	// user := User{Name: "sample", Email: "sample@sample.jp"}
	// db.Create(&user)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	defer db.Close()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello from go-docker-mysql-example",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		users := []User{}
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func connectDB() *gorm.DB {
	DBMS := os.Getenv("DB")
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME := os.Getenv("MYSQL_DATABASE")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("success db connection!!!")

	return db
}
