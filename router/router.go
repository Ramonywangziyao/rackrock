package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	var router = gin.New()

	var api = router.Group("/loginapi")
	{
		InitLoginRouter(api)
		InitBusinessRouter(api)
	}

	return router
}
