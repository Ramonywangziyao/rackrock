package controller

import (
	"github.com/gin-gonic/gin"
	"rackrock/context"
	"rackrock/model"
	"rackrock/service"
)

type DashboardController struct {
	BaseController
}

func (con DashboardController) GetBasic(c *gin.Context) model.RockResp {
	loginUser := context.GetLoginUser(c)
	dashboardBasic, err := service.GetDashboardInfo(loginUser.ID)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, model.SqlQueryError)
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: model.SqlQueryError,
			Data:    nil,
		}
	}

	con.Success(c, model.RequestSuccessMsg, dashboardBasic)
	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    dashboardBasic,
	}
}
