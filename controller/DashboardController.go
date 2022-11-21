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

func (con DashboardController) GetBasic(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	dashboardBasic, err := service.GetDashboardInfo(loginUser.ID)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, model.RequestSuccessMsg, dashboardBasic)
	return
}
