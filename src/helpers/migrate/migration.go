package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	db "github.com/saefullohmaslul/golang-example/src/database"
	"github.com/saefullohmaslul/golang-example/src/database/migration"
)

func main() {
	db.AppConnection()
	conn := db.GetDB()
	defer conn.Close()

	migration.CreateUser(conn)
}
