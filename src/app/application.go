package app

import (
	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/golang-example/src/database"
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
}

func configureAppDB() {
	database.AppConnection()
	db := database.GetDB()
	migration.Migrate(db)
}

func configureTestDB() {
	database.TestConnection()
	db := database.GetDB()
	migration.Migrate(db)
}
