package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/jpoles1/gopherbadger/logging"
	"github.com/saefullohmaslul/golang-example/src/app"
)

// main project
func main() {
	r := gin.Default()
	app := new(app.Application)
	app.CreateApp(r)

	if err := godotenv.Load(); err != nil {
		logging.Error("ENV", err)
	}

	port, found := os.LookupEnv("APP_PORT")
	if !found {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		logging.Error("APP", err)
	}
}
