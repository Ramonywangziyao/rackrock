package controller

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wenzhenxi/gorsa"
	"rackrock/context"
	"rackrock/model"
	"rackrock/service"
	"strings"
	"time"
)

type UserController struct {
	BaseController
}

func (con UserController) Login(c *gin.Context) model.RockResp {
	var loginRequest model.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
		return model.RockResp{
			Code:    model.RequestBodyErrorCode,
			Message: model.RequestBodyError,
			Data:    nil,
		}
	}

	// password encode
	var decode, err = gorsa.PriKeyDecrypt(loginRequest.Password, model.Pirvatekey)

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

	decodedData, err := hex.DecodeString(decode)

	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Decode Password Error. %s", err.Error()))
		con.Error(c, model.DataTypeConversionErrorCode, fmt.Sprintf("%s : %s", model.DataTypeConversionError, "password"))
		return model.RockResp{
			Code:    model.DataTypeConversionErrorCode,
			Message: fmt.Sprintf("%s : %s", model.DataTypeConversionError, "password"),
			Data:    nil,
		}
	}
	if strings.TrimSpace(queriedUser.Password) == fmt.Sprintf("%s", decodedData) {
		c.Set(context.LoginUser, model.LoginAccount{
			ID:       queriedUser.Id,
			UserName: queriedUser.Account,
		})

		err = service.SetLoginTime(queriedUser.Id)
		if err != nil {
			fmt.Println("c")
			fmt.Sprintf("更新登录时间错误")
		}

		loginResp := model.LoginResponse{}
		loginResp.Account = queriedUser.Account
		loginResp.LoginTime = time.Now()
		loginResp.LoginIp = c.ClientIP()
		loginResp.Token = context.CreateToken(queriedUser.Id, queriedUser.Account)

		con.Success(c, model.RequestSuccessMsg, loginResp)
		return model.RockResp{
			Code:    model.OK,
			Message: model.RequestSuccessMsg,
			Data:    loginResp,
		}
	}

	con.Error(c, model.PasswordErrorCode, model.PasswordError)
	return model.RockResp{
		Code:    model.PasswordErrorCode,
		Message: model.PasswordError,
		Data:    nil,
	}
}

func (con UserController) Register(c *gin.Context) model.RockResp {
	var registerRequest model.RegisterRequest
	if err := c.ShouldBind(&registerRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
		return model.RockResp{
			Code:    model.RequestBodyErrorCode,
			Message: model.RequestBodyError,
			Data:    nil,
		}
	}

	// 检查邀请码
	invitationCode := registerRequest.InvitationCode
	if invitationCode != model.InvitationCode {
		fmt.Println("not equal")
		con.Error(c, model.InvitationCodeErrorCode, model.InvitationCodeError)
		return model.RockResp{
			Code:    model.InvitationCodeErrorCode,
			Message: model.InvitationCodeError,
			Data:    nil,
		}
	}

	// 检查用户是否存在
	_, err := service.GetUserByAccount(registerRequest.Account)
	if err == nil {
		con.Error(c, model.RecordExistErrorCode, model.RecordExistError)
		return model.RockResp{
			Code:    model.RecordExistErrorCode,
			Message: model.RecordExistError,
			Data:    nil,
		}
	}

	// 创建用户
	id, err := service.CreateUser(registerRequest)
	if err != nil {
		con.Error(c, model.RegisterErrorCode, model.RegisterError)
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
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return model.RockResp{
			Code:    model.NotAuthorizedErrorCode,
			Message: model.NotAuthorizedError,
			Data:    nil,
		}
	}

	userListResponse, err := service.GetUserListResponse()
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : %s", model.SqlQueryError, "user list"))
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : %s", model.SqlQueryError, "user list"),
			Data:    nil,
		}
	}

	con.Success(c, model.RequestSuccessMsg, userListResponse)
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
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : %s", model.SqlQueryError, "user detail"))
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : %s", model.SqlQueryError, "user detail"),
			Data:    nil,
		}
	}

	con.Success(c, model.RequestSuccessMsg, user)
	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    user,
	}
}
