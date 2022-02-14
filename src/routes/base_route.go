package routes

import "go.uber.org/fx"

type Route interface {
	Setup()
}

var Module = fx.Options(
	fx.Provide(NewAccountRouter),
	fx.Provide(NewRoutes),
)

type Routes []Route

func NewRoutes(
	accountRouter AccountRouter,
) Routes {
	return Routes{
		&accountRouter,
	}
}

func (r *Routes) Setup() {
	for _, route := range *r {
		route.Setup()
	}
}
