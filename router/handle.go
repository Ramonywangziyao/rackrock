package router

import (
	"github.com/gin-gonic/gin"
	"rackrock/context"
	"rackrock/model"
	"time"
)

type ProcessHandle func(ctx *gin.Context)

func (handle ProcessHandle) Need(flag bool) {

}

type ServiceHandle func(ctx *gin.Context) model.RockResp

func Api(handler ServiceHandle) ProcessHandle {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

			}
			context.OperateHandler(ctx, context.AfterHandler)
		}()

		if err := context.OperateHandler(ctx, context.BeforeHandler); err != nil {
			panic(err)
		}

		var start = time.Now()
		var resp = handler(ctx)
		var dur = int64(time.Since(start).Seconds())
		ctx.Set(context.Response, resp)
		ctx.Set(context.Duration, dur)
	}
}
