package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rackrock/context"
	"rackrock/model"
	"rackrock/service"
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
	loginUser := context.GetLoginUser(c)
	accessLevel, err := service.GetUserAccessLevel(loginUser.ID)
	if err != nil {
		con.Error(c, fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"))
		return
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		con.Error(c, model.NotAuthorizedError)
		return
	}

	userListResponse, err := service.GetUserListResponse()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, model.RequestSuccessMsg, userListResponse)
	return
}

func (con UserController) UserDetail(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	userId := loginUser.ID

	user, err := service.GetUserDetail(userId)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, model.RequestSuccessMsg, user)
	return
}
