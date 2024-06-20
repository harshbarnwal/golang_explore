package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
)

func init() {

	var err error

	db, err = gorm.Open("mysql", "root:jvt123@tcp(127.0.0.1:3306)/sys?charset=utf8&parseTime=True")

	if err != nil {

	}

}

var db *gorm.DB

type BitCoin struct {
	gorm.Model

	Name  string
	Count int
}

func main() {
	ret := gin.Default()

	ret.GET("", func(c *gin.Context) {
		c.JSON(200, "Hello from go")
	})

	ret.POST("", func(c *gin.Context) {

		var trading BitCoin = BitCoin{"coindcx", 200}

		db.Create(&trading) // RAM to harddisk database

		c.JSON(200, "hello trading guys")

	})

	ret.Run()
}
