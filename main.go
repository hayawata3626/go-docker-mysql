package main

import (
  "fmt"
	"os"
	"log"

  "github.com/joho/godotenv"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load()
	r := gin.Default()
	db := connectDB()
	defer db.Close()

	if err != nil {
    log.Fatal("Error loading .env file")
	}


	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello from go-docker-mysql-example",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


func connectDB() * gorm.DB {
	DBMS     := os.Getenv("DB")
  USER     := os.Getenv("MYSQL_USER")
  PASS     := os.Getenv("MYSQL_PASSWORD")
  PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME   := os.Getenv("MYSQL_DATABASE")

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
    panic(err.Error())
	}

	fmt.Println("success db connection!!!")

  return db
}
