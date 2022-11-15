package rackrock

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"rackrock/utils"
)

var (
	Db *sql.DB
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Middleware executed.")
		c.Next()
	}
}

func Setup() *gin.Engine {
	Db = utils.DBConnect()
	r := gin.Default()
	r.Use(Middleware())
	return r
}
