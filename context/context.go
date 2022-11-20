package context

import (
	"context"
	"github.com/gin-gonic/gin"
	"rackrock/model"
)

func SetKV(ctx context.Context, key, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

const (
	LoginUser string = "login-account"
	Response  string = "response"
	Duration  string = "duration"
	IsAuth    string = "is-auth"
	Gin       string = "gin"
)

func GetLoginUser(ctx context.Context) model.LoginAccount {
	return ctx.Value(LoginUser).(model.LoginAccount)
}

func GetGinCtx(ctx context.Context) *gin.Context {
	return ctx.Value(Gin).(*gin.Context)
}

func GetResponse(ctx context.Context) model.RockResp {
	return ctx.Value(Response).(model.RockResp)
}

func IsNeedAuth(ctx context.Context) bool {
	if ctx.Value(IsAuth) == nil {
		return false
	}
	return ctx.Value(IsAuth).(bool)
}

func GetDuration(ctx context.Context) uint64 {
	return ctx.Value(Duration).(uint64)
}
