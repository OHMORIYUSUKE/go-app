package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "producer"
	PASS := "producer_pass"
	PROTOCOL := "tcp"
	HOST := "db"
	PORT := "3306"
	DBNAME := "idolm@ster"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "(" + HOST + ":" + PORT + ")" + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		db := gormConnect()
		log.Println(&db)
			c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
