package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	var router = gin.New()

	var api = router.Group("/api")
	{
		InitLoginRouter(api)
		InitBusinessRouter(api)
	}

	return router
}
