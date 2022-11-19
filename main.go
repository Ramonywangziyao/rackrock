package rackrock

import "rackrock/controller"

func main() {
	r := Setup()

	// routing paths here
	// general management
	genRouter := r.Group("general")
	generalController := controller.GeneralController{}
	genRouter.POST("/brand/creation", generalController.CreateBrand)
	genRouter.GET("/brand/list", generalController.GetBrandList)

	genRouter.POST("/tag/creation", generalController.CreateTag)
	genRouter.GET("/tag/list", generalController.GetTagList)

	genRouter.GET("/cities", controller.GeneralController{}.GetCities)

	// dashboard management
	dsbdRouter := r.Group("/dashboard")
	dashboardController := controller.DashboardController{}
	dsbdRouter.GET("/basics", dashboardController.GetBasic)

	// user management
	userRouter := r.Group("/user")
	userController := controller.UserController{}
	userRouter.POST("/login", userController.Login)
	userRouter.POST("/registration", userController.Register)
	userRouter.POST("/logout", userController.Logout)
	userRouter.GET("/list", userController.UserList)
	userRouter.GET("/detail", userController.UserDetail)

	// event management
	eventRouter := r.Group("/event")
	eventController := controller.EventController{}
	eventRouter.POST("/creation", eventController.CreateEvent)
	eventRouter.POST("/items", eventController.ImportItems)
	eventRouter.POST("/sold", eventController.ImportSold)
	eventRouter.POST("/return", eventController.ImportReturn)
	eventRouter.GET("/list", eventController.GetEventList)

	// member management
	memberRouter := r.Group("/member")
	memberController := controller.MemberController{}
	memberRouter.POST("/import", memberController.ImportMemberInfo)

	// report management
	reportRouter := r.Group("/report")
	reportController := controller.ReportController{}
	reportRouter.GET("/basic", reportController.GetBasic)
	reportRouter.GET("/shareLink", reportController.GetShareLink)
	reportRouter.GET("/ranking", reportController.GetShareLink)
	reportRouter.GET("/dailyDetail", reportController.GetDailyDetail)
	reportRouter.GET("/saleExport", reportController.ExportSaleDetail)

	r.Run(":3001")
}
