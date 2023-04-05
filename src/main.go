package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Member struct {
	Id                  int    `gorm:"" json:"Id"`
	Types               string `gorm:"" json:"Types"`
	Ch_name             string `gorm:"" json:"Ch_name"`
	Ch_name_ruby        string `gorm:"" json:"Ch_name_ruby"`
	Ch_family_name      string `gorm:"" json:"Ch_family_name"`
	Ch_family_name_ruby string `gorm:"" json:"Ch_family_name_ruby"`
	Ch_first_name       string `gorm:"" json:"Ch_first_name"`
	Ch_first_name_ruby  string `gorm:"" json:"Ch_first_name_ruby"`
	Ch_birth_month      int64  `gorm:"" json:"Ch_birth_month"`
	Ch_birth_day        int64  `gorm:"" json:"Ch_birth_day"`
	Ch_gender           int64  `gorm:"" json:"Ch_gender"`
	Is_idol             int64  `gorm:"" json:"Is_idol"`
	Ch_blood_types      string `gorm:"" json:"Ch_blood_types"`
	Ch_color            string `gorm:"" json:"Ch_color"`
	Cv_name             string `gorm:"" json:"userName"`
	Cv_name_ruby        string `gorm:"" json:"Cv_name_ruby"`
	Cv_family_name      string `gorm:"" json:"Cv_family_name"`
	Cv_family_name_ruby string `gorm:"" json:"Cv_family_name_ruby"`
	Cv_first_name       string `gorm:"" json:"Cv_first_name"`
	Cv_first_name_ruby  string `gorm:"" json:"Cv_first_name_ruby"`
	Cv_birth_month      int64  `gorm:"" json:"Cv_birth_month"`
	Cv_birth_day        int64  `gorm:"" json:"Cv_birth_day"`
	Cv_gender           int64  `gorm:"" json:"Cv_gender"`
	Cv_nickname         string `gorm:"" json:"Cv_nickname"`
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
	r.GET("/members", func(c *gin.Context) {
		db := gormConnect()
		var members []Member
		db.Debug().Table("imas_characters").Select("*").Scan(&members)
		for i := 0; i < len(members); i++ {
			log.Println(i)
			log.Println(members[i])
		}
		c.JSON(200, gin.H{
			"members": members,
		})
	})
	r.GET("/member", func(c *gin.Context) {
		name := c.Query("name")
		db := gormConnect()
		var member Member
		db.Debug().Table("imas_characters").Select("*").Where("ch_name LIKE ?", "%"+name+"%").First(&member)
		c.JSON(200, gin.H{
			"member": member,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
