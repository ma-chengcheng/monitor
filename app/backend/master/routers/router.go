package routers

import (
	"github.com/gin-gonic/gin"
	"master/controllers"
)

func InitRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	v1.GET("/ws/query", controllers.Query(controllers.WsConnHandle))
	v1.POST("/add_worker", controllers.AddWorker)
}