package controller

import (
	"github.com/gin-gonic/gin"
	"rackrock/model"
	"rackrock/service"
	"rackrock/utils"
)

type DashboardController struct {
	BaseController
}

func (con DashboardController) GetBasic(c *gin.Context) {
	userIdStr := c.Query("userId")
	userId, err := utils.ConvertStringToInt64(userIdStr)
	if err != nil {
		con.Error(c, model.RequestParameterError)
	}

	dashboardBasic, err := service.GetDashboardInfo(userId)
	if err != nil {
		con.Error(c, err.Error())
	}

	con.Success(c, model.RequestSuccessMsg, dashboardBasic)
}
