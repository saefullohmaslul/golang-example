package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/saefullohmaslul/Golang-Example/app"
	"github.com/saefullohmaslul/Golang-Example/database"
	"github.com/saefullohmaslul/Golang-Example/database/migration"
)

// main project
func main() {
	r := gin.Default()
	app := new(app.Application)
	app.CreateApp(r)

	database.Connection()
	db := database.GetDB()
	migration.Migrate(db)

	port := os.Getenv("APP_PORT")
	r.Run(":" + port)
}
