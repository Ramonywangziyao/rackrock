package router

import (
	"github.com/gin-gonic/gin"
	"rackrock/context"
	"rackrock/controller"
	"rackrock/loginapi"
)

func InitLoginRouter(router *gin.RouterGroup) {
	var userRouter = router.Group("user")
	var accountCon = loginapi.NewAccountCon()

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

	userController := controller.UserController{}
	userRouter.POST("/registration", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, false)
		Api(userController.Register)(ctx)
	})
	userRouter.GET("/list", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(userController.UserList)(ctx)
	})
	userRouter.GET("/detail", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(userController.UserDetail)(ctx)
	})
}
