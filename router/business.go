package router

import (
	"github.com/gin-gonic/gin"
	"rackrock/controller"
)

func InitBusinessRouter(router *gin.RouterGroup) {
	// routing paths here
	// general management
	genRouter := router.Group("general")
	generalController := controller.GeneralController{}
	genRouter.POST("/brand/creation", generalController.CreateBrand)
	genRouter.GET("/brand/list", generalController.GetBrandList)

	genRouter.POST("/tag/creation", generalController.CreateTag)
	genRouter.GET("/tag/list", generalController.GetTagList)

	genRouter.GET("/cities", controller.GeneralController{}.GetCities)

	// dashboard management
	dsbdRouter := router.Group("/dashboard")
	dashboardController := controller.DashboardController{}
	dsbdRouter.GET("/basics", dashboardController.GetBasic)

	// event management
	eventRouter := router.Group("/event")
	eventController := controller.EventController{}
	eventRouter.POST("/creation", eventController.CreateEvent)
	eventRouter.POST("/items", eventController.ImportItems)
	eventRouter.POST("/sold", eventController.ImportSold)
	eventRouter.POST("/return", eventController.ImportReturn)
	eventRouter.GET("/list", eventController.GetEventList)

	// member management
	memberRouter := router.Group("/member")
	memberController := controller.MemberController{}
	memberRouter.POST("/import", memberController.ImportMemberInfo)

	// report management
	reportRouter := router.Group("/report")
	reportController := controller.ReportController{}
	reportRouter.GET("/basic", reportController.GetBasic)
	reportRouter.GET("/shareLink", reportController.GetShareLink)
	reportRouter.GET("/ranking", reportController.GetShareLink)
	reportRouter.GET("/dailyDetail", reportController.GetDailyDetail)
	reportRouter.GET("/saleExport", reportController.ExportSaleDetail)
}
