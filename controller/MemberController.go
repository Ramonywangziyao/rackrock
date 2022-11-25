package controller

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"rackrock/context"
	"rackrock/model"
	"rackrock/service"
)

type MemberController struct {
	BaseController
}

func (con MemberController) ImportMember(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	accessLevel, err := service.GetUserAccessLevel(loginUser.ID)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"))
		return
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return
	}

	var importReturnRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importReturnRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		con.Error(c, model.ImportFileErrorCode, model.ImportFileError)
		return
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		con.Error(c, model.ExcelParseErrorCode, model.ExcelParseError)
		return
	}

	go service.ReadMember(xlsx)

	con.Success(c, model.RequestSuccessMsg, nil)
	return
}
