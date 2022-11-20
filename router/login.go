package router

import (
	"github.com/gin-gonic/gin"
	"rackrock/api"
	"rackrock/context"
)

func InitLoginRouter(router *gin.RouterGroup) {
	var userRouter = router.Group("user")
	var accountCon = api.NewAccountCon()

	// login
	userRouter.POST("login", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, false)
		Api(accountCon.Login)(ctx)
	})

	// test
	userRouter.GET("do-run", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, false)
		Api(accountCon.Query)(ctx)
	})

	//  register
	userRouter.POST("register", func(ctx *gin.Context) {

	})

}
