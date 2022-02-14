package main

import (
	"context"
	"fmt"
	"os"
	"restapi/src/controllers"
	"restapi/src/lib"
	"restapi/src/middlewares"
	"restapi/src/repositories"
	"restapi/src/routes"
	"restapi/src/services"
	"restapi/src/utils/databases"
	"time"

	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var Module = fx.Options(
	routes.Module,
	controllers.Module,
	services.Module,
	repositories.Module,
	databases.Module,
	lib.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	echoHandler lib.EchoHandler,
	routes routes.Routes,
	database databases.Database,
) {
	conn, _ := database.DB.DB()
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				conn.SetMaxIdleConns(10)
				conn.SetMaxOpenConns(100)
				conn.SetConnMaxLifetime(time.Hour)

				go func() {
					port, found := os.LookupEnv("PORT")

					if !found {
						port = "1323"
					}

					echoHandler.Echo.Validator = middlewares.NewValidator(validator.New())
					echoHandler.Echo.HTTPErrorHandler = middlewares.ErrorHandler

					routes.Setup()

					echoHandler.Echo.Logger.Fatal(
						echoHandler.Echo.Start(fmt.Sprintf(":%s", port)),
					)
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				return conn.Close()
			},
		},
	)
}
