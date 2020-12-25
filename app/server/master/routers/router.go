package routers

import (
	"github.com/gin-gonic/gin"
	"master/controllers"
	"net/http"
)

func InitRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	router.LoadHTMLFiles("/src/www/index.html")
	v1.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	v1.POST("/login", controllers.Login)
	v1.POST("/register", controllers.Register)

	//v1.Use(middlewares.JWT())
	//{
	v1.GET("/ws/query", controllers.QueryNodesAPI(controllers.WsConnHandle))
	v1.POST("/add_node", controllers.AddNodeAPI)
	v1.POST("/delete_node", controllers.DeleteNodeAPI)
	//}
}
