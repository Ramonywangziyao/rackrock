package controller

import (
	"fmt"
	"github.com/farmerx/gorsa"
	"github.com/gin-gonic/gin"
	"rackrock/context"
	"rackrock/model"
	"rackrock/service"
	"strings"
	"time"
)

type UserController struct {
	BaseController
}

func (con UserController) Login(c *gin.Context) (res model.RockResp) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
		return
	}

	// password encode
	var decode, err = gorsa.RSA.PriKeyDECRYPT([]byte(loginRequest.Password))
	var account = loginRequest.Account
	// 需要改这个方法，拿到用户密码等信息
	queriedUser, err := service.GetUserByAccount(account)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : %s", model.SqlQueryError, "user"))
		return
	}

	if strings.TrimSpace(queriedUser.Password) == string(decode) {
		c.Set(context.LoginUser, model.LoginAccount{
			ID:       queriedUser.Id,
			UserName: queriedUser.Account,
		})

		err = service.SetLoginTime(queriedUser.Id)
		if err != nil {
			fmt.Sprintf("更新登录时间错误")
		}

		con.Success(c, model.RequestSuccessMsg, map[string]interface{}{
			"username":  queriedUser.Account,
			"loginIp":   c.ClientIP(),
			"loginTime": time.Now(),
			"token":     context.CreateToken(queriedUser.Id, queriedUser.Account),
		})
		return
	}

	con.Error(c, model.PasswordErrorCode, model.PasswordError)
	return
}

func (con UserController) Register(c *gin.Context) (res model.RockResp) {
	var registerRequest model.RegisterRequest
	if err := c.ShouldBind(&registerRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
		return
	}

	// 检查邀请码
	invitationCode := registerRequest.InvitationCode
	if invitationCode != model.InvitationCode {
		con.Error(c, model.InvitationCodeErrorCode, model.InvitationCodeError)
		return
	}

	// 检查用户是否存在
	_, err := service.GetUserByAccount(registerRequest.Account)
	if err == nil {
		con.Error(c, model.RecordExistErrorCode, model.RecordExistError)
		return
	}

	// 创建用户
	id, err := service.CreateUser(registerRequest)
	if err != nil {
		con.Error(c, model.RegisterErrorCode, model.RegisterError)
		return
	}

	con.Success(c, model.RequestSuccessMsg, id)
	return
}

func (con UserController) UserList(c *gin.Context) (res model.RockResp) {
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

	userListResponse, err := service.GetUserListResponse()
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : %s", model.SqlQueryError, "user list"))
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
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : %s", model.SqlQueryError, "user detail"))
		return
	}

	con.Success(c, model.RequestSuccessMsg, user)
	return
}
