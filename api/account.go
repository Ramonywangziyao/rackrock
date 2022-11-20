package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"rackrock/context"
	"rackrock/model"
	"rackrock/repository/domain"
	"rackrock/service"
	"rackrock/utils"
	"strings"
	"time"
)

type AccountController struct {
	Account service.Account
}

func NewAccountCon() AccountController {
	return AccountController{Account: service.GetAccountService()}
}

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (self *AccountController) Query(ctx *gin.Context) (res model.RockResp) {
	var username = ctx.Query("username")

	var account = &domain.Account{Username: username}
	var err = self.Account.GetAccount(account, "password", "id")
	if err != nil {
		utils.NewBusinessErr("account not found")
	}

	ctx.Set(context.LoginUser, model.LoginAccount{
		ID:       account.Id,
		UserName: account.Username,
	})
	//ctx = context.SetKV(ctx, context.LoginUser, model.LoginAccount{
	//	ID:       account.Id,
	//	UserName: account.Username,
	//})

	res = model.RockResp{
		Code:    0,
		Message: "success",
		Data:    *account,
	}

	ctx.JSON(200, res)
	return
}

func (self *AccountController) Login(ctx *gin.Context) model.RockResp {

	var loginForm = new(LoginForm)
	BindJson(ctx, loginForm)

	// password encode
	var encode = loginForm.Password

	var account = &domain.Account{
		Username: loginForm.Username,
	}
	var err = self.Account.GetAccount(account, "password", "status", "id")
	if err != nil {
		utils.NewBusinessErr("account not found")
	}

	if strings.TrimSpace(account.Password) == encode {
		ctx.Set(context.LoginUser, model.LoginAccount{
			ID:       account.Id,
			UserName: account.Username,
		})
		return model.RockResp{
			Code:    0,
			Message: "success",
			Data: map[string]interface{}{
				"username":  account.Username,
				"loginIp":   ctx.ClientIP(),
				"loginTime": time.Now(),
				"token":     context.CreateToken(account.Id, account.Username),
			},
		}
	} else {
		err = errors.New("password err")
	}

	return model.RockResp{
		Code:    utils.ServiceError.Code,
		Message: err.Error(),
		Data:    nil,
	}
}
