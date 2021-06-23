package router

import (
	"../controller"

	"github.com/gin-gonic/gin"
)

func SetUpRoute() *gin.Engine {
	r := gin.Default()

	//r.Static("/static", "static")
	//r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	//v1Group := r.Group("v1")
	//{
	//	v1Group.POST("/todo", controller.CreateATodo)
	//	v1Group.GET("/todo", controller.GetTodoList)
	//	v1Group.PUT("/todo/:id", controller.UpdateATodo)
	//	v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	//}
	return r
}
