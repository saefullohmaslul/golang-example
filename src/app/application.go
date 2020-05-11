package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/saefullohmaslul/golang-example/src/database"
	"github.com/saefullohmaslul/golang-example/src/database/migration"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
	"github.com/saefullohmaslul/golang-example/src/routes"
	"github.com/saefullohmaslul/golang-example/src/runner"
)

// Application -> application instance
type Application struct {
}

// CreateApp -> method to create gin application
func (a Application) CreateApp(r *gin.Engine) {
	r.Use(exception.Recovery(exception.ErrorHandler))
	configureAPIEndpoint(r)
	configureAppDB()
	runner.PubSubRunner()
}

// CreateTest -> method to create gin application with environment test
func (a Application) CreateTest(r *gin.Engine) {
	r.Use(exception.Recovery(exception.ErrorHandler))
	configureAPIEndpoint(r)
	configureTestDB()
	runner.PubSubRunner()
}

/**
 * configuration all endpoint
 */
func configureAPIEndpoint(r *gin.Engine) {
	g := r.Group("/")
	routes.Router(g)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"flag":   "NOT_FOUND",
			"errors": gin.H{
				"message": "The route you are looking for is not found",
				"flag":    "ROUTE_NOT_FOUND",
			},
		})
	})
}

/**
 * configuration database application
 */
func configureAppDB() {
	db.AppConnection()
	conn := db.GetDB()
	migration.Migrate(conn)
}

/**
 * configuration database application for testing
 */
func configureTestDB() {
	db.TestConnection()
	conn := db.GetDB()
	migration.Migrate(conn)
}
