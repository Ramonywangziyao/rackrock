package rackrock

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rackrock/setting"
	"rackrock/utils"
)

var (
	Db *gorm.DB
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Middleware executed.")
		c.Next()
	}
}

func Setup() *gin.Engine {
	Db = utils.DBConnect()
	setting.DB = Db
	r := gin.Default()
	r.Use(Middleware())
	return r
}
