package router

import (
	"github.com/gin-gonic/gin"
)

func InitBusinessRouter(router *gin.RouterGroup) {
	var rockRouter = router.Group("rock")

	rockRouter.GET("/user", func(context *gin.Context) {

	})
}
