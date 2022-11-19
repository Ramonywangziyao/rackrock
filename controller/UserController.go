package controller

import (
	"github.com/gin-gonic/gin"
	"rackrock/model"
)

type UserController struct {
	BaseController
}

func (con UserController) Login(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}
}

func (con UserController) Register(c *gin.Context) {
	var registerRequest model.RegisterRequest
	if err := c.ShouldBind(&registerRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}
}

func (con UserController) Logout(c *gin.Context) {
	var logoutRequest model.LogOutRequest
	if err := c.ShouldBind(&logoutRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}
}

func (con UserController) UserList(c *gin.Context) {

}

func (con UserController) UserDetail(c *gin.Context) {

}
