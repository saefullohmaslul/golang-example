package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/golang-example/src/controllers"
	"github.com/saefullohmaslul/golang-example/src/validation"
)

// Router middleware to handler routes
func Router(g *gin.RouterGroup) {
	controller := controllers.UserController{}
	{
		g.GET("/users", controller.GetUsers)
		g.GET("/user/:id", validation.GetUser, controller.GetUser)
		g.POST("/user", validation.CreateUser, controller.CreateUser)
	}
}
