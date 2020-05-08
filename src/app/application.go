package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/saefullohmaslul/golang-example/src/database"
	"github.com/saefullohmaslul/golang-example/src/database/migration"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
	"github.com/saefullohmaslul/golang-example/src/routes"
)

// Application struct
type Application struct {
}

// CreateApp method
func (a Application) CreateApp(r *gin.Engine) {
	r.Use(exception.Recovery(exception.ErrorHandler))
	configureAPIEndpoint(r)
	configureAppDB()
}

// CreateTest method
func (a Application) CreateTest(r *gin.Engine) {
	r.Use(exception.Recovery(exception.ErrorHandler))
	configureAPIEndpoint(r)
	configureTestDB()
}

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

func configureAppDB() {
	db.AppConnection()
	conn := db.GetDB()
	migration.Migrate(conn)
}

func configureTestDB() {
	db.TestConnection()
	conn := db.GetDB()
	migration.Migrate(conn)
}
