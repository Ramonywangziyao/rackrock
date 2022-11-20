package api

import (
	"github.com/gin-gonic/gin"
	"rackrock/utils"
)

func BindJson(ctx *gin.Context, data interface{}) {
	if err := ctx.ShouldBindJSON(data); err != nil {
		panic(utils.NewBusinessErr(err.Error()))
	}
}

func BindQuery(ctx *gin.Context, data interface{}) {
	if err := ctx.ShouldBindQuery(data); err != nil {
		panic(utils.NewBusinessErr(err.Error()))
	}
}
