package app

import (
	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/Golang-Example/middlewares/exception"
	"github.com/saefullohmaslul/Golang-Example/routes"
)

// Application struct
type Application struct {
}

// CreateApp method
func (a Application) CreateApp(r *gin.Engine) {
	r.Use(exception.Recovery(exception.ErrorHandler))
	configureAPIEndpoint(r)
}

func configureAPIEndpoint(r *gin.Engine) {
	g := r.Group("/user")
	routes.Router(g)
}
