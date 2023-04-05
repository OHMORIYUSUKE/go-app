package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type Member struct {
    id int `gorm:"" json:"userName"`
	types string `gorm:"" json:"userName"`
    ch_name string `gorm:"" json:"userName"`
	ch_name_ruby string `gorm:"" json:"userName"`
	ch_family_name string `gorm:"" json:"userName"`
	ch_family_name_ruby string `gorm:"" json:"userName"`
	ch_first_name string `gorm:"" json:"userName"`
	ch_first_name_ruby string `gorm:"" json:"userName"`
	ch_birth_month int64 `gorm:"" json:"userName"`
	ch_birth_day int64 `gorm:"" json:"userName"`
	ch_gender int64 `gorm:"" json:"userName"`
	is_idol int64 `gorm:"" json:"userName"`
	ch_blood_types string `gorm:"" json:"userName"`
	ch_color string `gorm:"" json:"userName"`
	cv_name string `gorm:"" json:"userName"`
	cv_name_ruby string `gorm:"" json:"userName"`
	cv_family_name string `gorm:"" json:"userName"`
	cv_family_name_ruby string `gorm:"" json:"userName"`
	cv_first_name string `gorm:"" json:"userName"`
	cv_first_name_ruby string `gorm:"" json:"userName"`
	cv_birth_month int64 `gorm:"" json:"userName"`
	cv_birth_day int64 `gorm:"" json:"userName"`
	cv_gender int64 `gorm:"" json:"userName"`
	cv_nickname string `gorm:"" json:"userName"`
}

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
		//
		user1 := Member{}
		db.Debug().Table("imas_characters").First(&user1)
		// SELECT * FROM users ORDER BY id LIMIT 1;
		fmt.Println("first:", user1)
		//
		// var members []Member
		// db.Debug().Table("imas_characters").Select("*").Scan(&members)
		// for i := 0; i < len(members); i++ {
		// 	log.Println(i)
		// 	log.Println(members[i])
		// }
		c.JSON(200, gin.H{
			"message":"a",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
