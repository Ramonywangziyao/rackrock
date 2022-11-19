package controller

import (
	"github.com/gin-gonic/gin"
	"rackrock/model"
)

type MemberController struct {
	BaseController
}

func (con MemberController) ImportMemberInfo(c *gin.Context) {
	var importMemberRequest model.ImportMemberRequest
	if err := c.ShouldBind(&importMemberRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}
}
