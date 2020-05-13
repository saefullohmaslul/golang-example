package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/jpoles1/gopherbadger/logging"
	"github.com/saefullohmaslul/golang-example/src/apps"
)

/**
 * Main project
 * Framework: go-gin framework
 * ORM: gorm
 * Author: Saefulloh Maslul
 * License: MIT
 */
func main() {
	r := gin.Default()

	/**
	 * create application instance
	 * and configure with gin engine
	 */
	app := new(apps.Application)
	app.CreateApp(r)

	/**
	 * load the default env
	 */
	if err := godotenv.Load(); err != nil {
		logging.Error("ENV", err)
	}

	/**
	 * setting up port application
	 */
	port, found := os.LookupEnv("APP_PORT")
	if !found {
		port = "8080"
	}

	/**
	 * running application
	 */
	if err := r.Run(":" + port); err != nil {
		logging.Error("APP", err)
	}
}
