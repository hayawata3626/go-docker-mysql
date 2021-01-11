package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gin_test/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	r := gin.Default()
	r.Use(cors.Default())
	db := connectDB()
	db.AutoMigrate(&model.User{})

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	defer db.Close()

	r.GET("/users", func(c *gin.Context) {
		users := []model.User{}
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	r.POST("/user/new", func(c *gin.Context) {
		var req model.User
		c.BindJSON(&req)
		db.Create(&model.User{Name: req.Name, Email: req.Email})
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}

		c.Redirect(302, "/")
	})

	r.DELETE("user/:id", func(c *gin.Context) {
		var user model.User
		id := c.Param("id")

		db.Where("ID = ?", id).Delete(&user)
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
