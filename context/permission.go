package context

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"rackrock/config"
	"rackrock/logger"
	"rackrock/model"
	"rackrock/utils"
	"time"
)

func PermissionHandle(ctx context.Context) (err error) {

	if !IsNeedAuth(ctx) {
		return nil
	}

	var ginCtx = ctx.(*gin.Context)

	var token = ginCtx.Request.Header.Get(utils.JwtKey)
	if utils.IsEmptyStr(token) {
		token = ginCtx.Query("token")
	}

	if utils.IsEmptyStr(token) {
		return utils.AuthError
	}

	var account = ParseToken(token)
	if account == nil {
		return utils.AuthError
	}

	ctx = SetKV(ctx, LoginUser, *account)
	return
}

func ParseToken(token string) *model.LoginAccount {
	if utils.IsEmptyStr(token) {
		return nil
	}

	var bytes, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.Jwt.Key), nil
	})

	if err != nil {
		logger.Logger.Error(fmt.Sprintf("parse token err: %s", err.Error()))
		return nil
	}

	var claims = bytes.Claims.(jwt.MapClaims)
	return &model.LoginAccount{
		ID:       uint64(claims["id"].(float64)),
		UserName: claims["user_name"].(string),
	}
}

func CreateToken(userId uint64, name string) string {
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        userId,
		"user_name": name,
		"expire":    time.Now().Add(time.Minute * time.Duration(config.Cfg.Jwt.ExpireTime)).Unix(),
	})
	k := []byte(config.Cfg.Jwt.Key)
	var str, err = token.SignedString(k)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %+s", err))
		panic(fmt.Sprintf("create token err: %s", err.Error()))
	}

	return str
}
