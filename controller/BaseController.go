package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct{}

func (con BaseController) Success(c *gin.Context, code int, msg string, data interface{}) {
	fmt.Println(fmt.Sprintf("data %+s", data))
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func (con BaseController) Error(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": errorCode,
		"msg":  msg,
	})
}
