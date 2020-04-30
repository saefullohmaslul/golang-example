package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/Golang-Example/controllers"
)

// Router middleware to handler routes
func Router(g *gin.RouterGroup) {
	controller := controllers.UserController{}
	{
		g.GET("/name", controller.GetName)
		g.GET("/biodata", controller.GetBiodata)
		g.POST("/create", controller.CreateUser)
	}
}
