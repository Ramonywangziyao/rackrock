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

func (con UserController) Login(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}

	con.Success(c, model.RequestSuccessMsg, nil)
}

func (con UserController) Register(c *gin.Context) {
	var registerRequest model.RegisterRequest
	if err := c.ShouldBind(&registerRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}

	con.Success(c, model.RequestSuccessMsg, nil)
}

func (con UserController) Logout(c *gin.Context) {
	var logoutRequest model.LogOutRequest
	if err := c.ShouldBind(&logoutRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}

	con.Success(c, model.RequestSuccessMsg, nil)
}

func (con UserController) UserList(c *gin.Context) {

}

func (con UserController) UserDetail(c *gin.Context) {
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

}
