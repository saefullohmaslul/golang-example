package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/golang-example/src/controllers"
	"github.com/saefullohmaslul/golang-example/src/validations"
)

// Router middleware to handler routes
func Router(g *gin.RouterGroup) {
	controller := controllers.UController()
	{
		g.GET("/users", controller.GetUsers)
		g.GET("/user/:id", validations.GetUser, controller.GetUser)
		g.POST("/user", validations.CreateUser, controller.CreateUser)
		g.PATCH("/user/:id", validations.UpdateUser, controller.UpdateUser)
		g.DELETE("/user/:id", validations.DeleteUser, controller.DeleteUser)

		g.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":   http.StatusOK,
				"message":  "ready",
				"database": "error",
			})
		})
	}
}
