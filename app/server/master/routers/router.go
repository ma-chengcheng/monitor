package routers

import (
	"github.com/gin-gonic/gin"
	"master/controllers"
)

func InitRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	v1.POST("/login", controllers.Login)
	v1.POST("/register", controllers.Register)

	//v1.Use(middlewares.JWT())
	//{
	v1.GET("/ws/query", controllers.Query(controllers.WsConnHandle))
	v1.POST("/add_node", controllers.AddNode)
	v1.POST("/delete_node", controllers.DeleteNode)
	//}
}
