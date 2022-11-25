package router

import (
	"github.com/gin-gonic/gin"
	"rackrock/context"
	"rackrock/controller"
)

func InitBusinessRouter(router *gin.RouterGroup) {
	// routing paths here
	// general management
	genRouter := router.Group("general")
	generalController := controller.GeneralController{}
	genRouter.POST("/brand/creation", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(generalController.CreateBrand)(ctx)
	})
	genRouter.GET("/brand/list", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, false)
		Api(generalController.GetBrandList)(ctx)
	})

	genRouter.POST("/tag/creation", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(generalController.CreateTag)(ctx)
	})
	genRouter.GET("/tag/list", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(generalController.GetTagList)(ctx)
	})

	genRouter.GET("/cities", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(controller.GeneralController{}.GetCities)(ctx)
	})

	genRouter.GET("/industryList", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(controller.GeneralController{}.GetIndustries)(ctx)
	})

	// dashboard management
	dsbdRouter := router.Group("/dashboard")
	dashboardController := controller.DashboardController{}
	dsbdRouter.GET("/basics", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(dashboardController.GetBasic)(ctx)
	})

	// event management
	eventRouter := router.Group("/event")
	eventController := controller.EventController{}
	eventRouter.POST("/creation", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(eventController.CreateEvent)(ctx)
	})
	eventRouter.POST("/items", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(eventController.ImportItems)(ctx)
	})
	eventRouter.POST("/sold", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(eventController.ImportSold)(ctx)
	})
	eventRouter.POST("/return", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(eventController.ImportReturn)(ctx)
	})
	eventRouter.GET("/list", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(eventController.GetEventList)(ctx)
	})

	// member management
	memberRouter := router.Group("/member")
	memberController := controller.MemberController{}
	memberRouter.POST("/import", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(memberController.ImportMember)(ctx)
	})

	// report management
	reportRouter := router.Group("/report")
	reportController := controller.ReportController{}
	reportRouter.GET("/basic", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(reportController.GetBasic)(ctx)
	})
	reportRouter.GET("/shareLink", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(reportController.GetShareLink)(ctx)
	})
	reportRouter.GET("/ranking", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(reportController.GetShareLink)(ctx)
	})
	reportRouter.GET("/dailyDetail", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(reportController.GetDailyDetail)(ctx)
	})
	reportRouter.GET("/saleExport", func(ctx *gin.Context) {
		ctx.Set(context.IsAuth, true)
		Api(reportController.ExportSaleDetail)(ctx)
	})
}
