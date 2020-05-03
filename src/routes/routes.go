package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/golang-example/src/controllers"
)

// Router middleware to handler routes
func Router(g *gin.RouterGroup) {
	controller := controllers.UserController{}
	{
		g.GET("/name", controller.GetName)
		g.GET("/biodata", controller.GetBiodata)
		g.POST("/user", controller.CreateUser)
	}
}
