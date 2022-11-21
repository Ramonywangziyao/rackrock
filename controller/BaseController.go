package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct{}

func (con BaseController) Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
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
