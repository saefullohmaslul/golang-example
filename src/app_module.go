package main

import (
	"context"
	"fmt"
	"os"
	"restapi/src/domains/accounts"
	"restapi/src/lib"
	"restapi/src/middlewares"
	"restapi/src/repositories"
	"time"

	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var Module = fx.Options(
	accounts.Module,
	repositories.Module,
	lib.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	echoHandler lib.EchoHandler,
	database lib.Database,
	accounts accounts.AccountRoute,
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

					accounts.Setup()

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
