package context

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"rackrock/config"
	"rackrock/logger"
	"rackrock/model"
	"rackrock/utils"
	"time"
)

type MyCustomClaims struct {
	ID       uint64 `json:"id"`
	UserName string `json:"user_name"`
	Expire   int64  `json:"expire"`
	jwt.StandardClaims
}

func PermissionHandle(ctx *gin.Context) error {
	if !IsNeedAuth(ctx) {
		return nil
	}

	var token = ctx.Request.Header.Get(utils.JwtKey)
	if utils.IsEmptyStr(token) {
		token = ctx.Query("token")
	}

	if utils.IsEmptyStr(token) {
		return utils.AuthError
	}

	var account = ParseToken(token)
	fmt.Println(fmt.Sprintf("Account: %+s", account))
	if account == nil {
		return utils.AuthError
	}

	//SetKV(ctx, LoginUser, *account)
	ctx.Set(LoginUser, *account)
	return nil
}

func ParseToken(token string) *model.LoginAccount {
	if utils.IsEmptyStr(token) {
		return nil
	}

	var bytes, err = jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.Jwt.Key), nil
	})

	if err != nil {
		logger.Logger.Error(fmt.Sprintf("parse token err: %s", err.Error()))
		return nil
	}

	var claims = bytes.Claims.(*MyCustomClaims)
	//fmt.Println(fmt.Sprintf("Id %s, err : %s", id, err))
	return &model.LoginAccount{
		ID:       claims.ID,
		UserName: claims.UserName,
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
