package interfaces

import (
	"go.uber.org/dig"
	"go.uber.org/zap"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/http/server"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/log"
)

type routesDependencies struct {
	dig.In
	Handlers *ExampleHandler `name:"EXAMPLE-HANDLER"`
	Logger   *log.Logger     `name:"LOGGER"`
}

func ExampleRoutes(container *dig.Container) []server.Route {
	var dep routesDependencies

	if err := container.Invoke(func(h routesDependencies) {
		dep = h
	}); err != nil {
		dep.Logger.Fatal("failed to invoke dependencies module EXAMPLE", zap.Error(err))
	}

	return []server.Route{
		{
			Method:      "GET",
			Path:        "/find",
			Handler:     dep.Handlers.Find,
			Description: "Find example",
			Middlewares: nil,
		},
	}
}

func middlewares() {

}
