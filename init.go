package rackrock

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
