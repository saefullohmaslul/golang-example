package main

import (
	"restapi/src/controllers"
	"restapi/src/repositories"
	"restapi/src/services"
	"restapi/src/utils/databases"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type Module struct {
}

func (m *Module) New(e *echo.Echo) {
	ioc := m.NewIOC()

	s := NewService(ioc)

	s.NewRoute(e)
}

func (m *Module) NewIOC() di.Container {
	builder, _ := di.NewBuilder()
	builder.Add(di.Def{
		Name: "database",
		Build: func(ctn di.Container) (interface{}, error) {
			db, err := databases.Create()
			return db, err
		},
	})

	builder.Add(di.Def{
		Name: "repository",
		Build: func(ctn di.Container) (interface{}, error) {
			return repositories.NewRepository(builder.Build()), nil
		},
	})

	builder.Add(di.Def{
		Name: "service",
		Build: func(ctn di.Container) (interface{}, error) {
			return services.NewService(builder.Build()), nil
		},
	})

	builder.Add(di.Def{
		Name: "controller",
		Build: func(ctn di.Container) (interface{}, error) {
			return controllers.NewController(builder.Build()), nil
		},
	})

	return builder.Build()
}
