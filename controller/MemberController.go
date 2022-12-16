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
		con.Error(c, model.SqlQueryErrorCode, model.SqlQueryError)
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"),
			Data:    nil,
		}
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return model.RockResp{
			Code:    model.NotAuthorizedErrorCode,
			Message: model.NotAuthorizedError,
			Data:    nil,
		}
	}

	var importReturnRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importReturnRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
		return model.RockResp{
			Code:    model.RequestBodyErrorCode,
			Message: model.RequestBodyError,
			Data:    nil,
		}
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		con.Error(c, model.ImportFileErrorCode, model.ImportFileError)
		return model.RockResp{
			Code:    model.ImportFileErrorCode,
			Message: model.ImportFileError,
			Data:    nil,
		}
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		con.Error(c, model.ExcelParseErrorCode, model.ExcelParseError)
		return model.RockResp{
			Code:    model.ExcelParseErrorCode,
			Message: model.ExcelParseError,
			Data:    nil,
		}
	}

	go service.ReadMember(xlsx)

	con.Success(c, 0, model.RequestSuccessMsg, nil)
	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    nil,
	}
}
