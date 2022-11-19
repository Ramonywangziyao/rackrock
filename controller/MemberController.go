package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rackrock/model"
	"rackrock/utils"
)

type MemberController struct {
	BaseController
}

func (con MemberController) ImportMemberInfo(c *gin.Context) {
	var importMemberRequest model.ImportMemberRequest
	if err := c.ShouldBind(&importMemberRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}

	c.JSON(http.StatusOK, utils.GetHttpResponse(model.OK, model.RequestSuccessMsg, nil))
}
