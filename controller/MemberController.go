package controller

import (
	"github.com/gin-gonic/gin"
	"rackrock/model"
)

type MemberController struct {
	BaseController
}

func (con MemberController) ImportMemberInfo(c *gin.Context) (res model.RockResp) {
	var importMemberRequest model.ImportMemberRequest
	if err := c.ShouldBind(&importMemberRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
		return
	}

	con.Success(c, model.RequestSuccessMsg, nil)
	return
}
