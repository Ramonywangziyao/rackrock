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
		return model.RockResp{
			Code:    model.RequestBodyErrorCode,
			Message: model.RequestBodyError,
			Data:    nil,
		}
	}

	// password encode
	var decode, err = gorsa.RSA.PriKeyDECRYPT([]byte(loginRequest.Password))
	var account = loginRequest.Account
	// 需要改这个方法，拿到用户密码等信息
	queriedUser, err := service.GetUserByAccount(account)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : %s", model.SqlQueryError, "user"))
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : %s", model.SqlQueryError, "user"),
			Data:    nil,
		}
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

		return model.RockResp{
			Code:    model.OK,
			Message: model.RequestSuccessMsg,
			Data: map[string]interface{}{
				"username":  queriedUser.Account,
				"loginIp":   c.ClientIP(),
				"loginTime": time.Now(),
				"token":     context.CreateToken(queriedUser.Id, queriedUser.Account),
			},
		}
	}

	return model.RockResp{
		Code:    model.PasswordErrorCode,
		Message: model.PasswordError,
		Data:    nil,
	}
}

func (con UserController) Register(c *gin.Context) (res model.RockResp) {
	var registerRequest model.RegisterRequest
	if err := c.ShouldBind(&registerRequest); err != nil {
		return model.RockResp{
			Code:    model.RequestBodyErrorCode,
			Message: model.RequestBodyError,
			Data:    nil,
		}
	}

	// 检查邀请码
	invitationCode := registerRequest.InvitationCode
	if invitationCode != model.InvitationCode {
		return model.RockResp{
			Code:    model.InvitationCodeErrorCode,
			Message: model.InvitationCodeError,
			Data:    nil,
		}
	}

	// 检查用户是否存在
	_, err := service.GetUserByAccount(registerRequest.Account)
	if err == nil {
		return model.RockResp{
			Code:    model.RecordExistErrorCode,
			Message: model.RecordExistError,
			Data:    nil,
		}
	}

	// 创建用户
	id, err := service.CreateUser(registerRequest)
	if err != nil {
		return model.RockResp{
			Code:    model.RegisterErrorCode,
			Message: model.RegisterError,
			Data:    nil,
		}
	}

	con.Success(c, model.RequestSuccessMsg, id)
	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    id,
	}
}

func (con UserController) UserList(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	accessLevel, err := service.GetUserAccessLevel(loginUser.ID)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"))
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"),
			Data:    nil,
		}
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		return model.RockResp{
			Code:    model.NotAuthorizedErrorCode,
			Message: model.NotAuthorizedError,
			Data:    nil,
		}
	}

	userListResponse, err := service.GetUserListResponse()
	if err != nil {
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : %s", model.SqlQueryError, "user list"),
			Data:    nil,
		}
	}

	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    userListResponse,
	}
}

func (con UserController) UserDetail(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	userId := loginUser.ID

	user, err := service.GetUserDetail(userId)
	if err != nil {
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : %s", model.SqlQueryError, "user detail"),
			Data:    nil,
		}
	}

	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    user,
	}
}
