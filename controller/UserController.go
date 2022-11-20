package controller

import (
	"github.com/gin-gonic/gin"
	"rackrock/model"
	"rackrock/service"
	"rackrock/utils"
)

type UserController struct {
	BaseController
}

func (con UserController) Login(c *gin.Context) (res model.RockResp) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		con.Error(c, model.RequestBodyError)
		return
	}

	con.Success(c, model.RequestSuccessMsg, nil)
	return
}

func (con UserController) Register(c *gin.Context) (res model.RockResp) {
	var registerRequest model.RegisterRequest
	if err := c.ShouldBind(&registerRequest); err != nil {
		con.Error(c, model.RequestBodyError)
		return
	}

	con.Success(c, model.RequestSuccessMsg, nil)
	return
}

func (con UserController) UserList(c *gin.Context) (res model.RockResp) {
	return
}

func (con UserController) UserDetail(c *gin.Context) (res model.RockResp) {
	userIdStr := c.Query("userId")
	userId, err := utils.ConvertStringToInt64(userIdStr)
	if err != nil {
		con.Error(c, model.RequestParameterError)
		return
	}

	user, err := service.GetUserDetail(userId)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, model.RequestSuccessMsg, user)
	return
}
