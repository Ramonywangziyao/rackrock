package router

import (
	"github.com/gin-gonic/gin"
	"rackrock/context"
	"rackrock/controller"
)

func InitLoginRouter(router *gin.RouterGroup) {
	var userRouter = router.Group("user")
	// var accountCon = loginapi.NewAccountCon()
	userController := controller.UserController{}
	// login
	userRouter.POST("login", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, false)
		Api(userController.Login)(ctx)
	})

	//// test
	//userRouter.GET("do-run", func(ctx *gin.Context) {
	//	ctx.Set(context.IsAuth, false)
	//	Api(accountCon.Query)(ctx)
	//})

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
