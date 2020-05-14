package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	db "github.com/saefullohmaslul/golang-example/src/database"
	"github.com/saefullohmaslul/golang-example/src/database/seed"
	"github.com/sirupsen/logrus"
)

func main() {
	db.SeedConnection()
	conn := db.GetDB()
	defer conn.Close()

	for _, seed := range seed.All() {
		if err := seed.Run(conn); err != nil {
			logrus.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}
