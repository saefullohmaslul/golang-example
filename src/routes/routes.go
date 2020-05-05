package routes

import (
	"net/http"

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
		g.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":   http.StatusOK,
				"message":  "ready",
				"database": "error",
			})
		})
	}
}
