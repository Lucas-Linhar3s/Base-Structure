package example

import (
	"go.uber.org/dig"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/modules/example/application"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/modules/example/domain"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/modules/example/interfaces"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/http/server"
)

type dependenciesModel struct {
	Constructor interface{}
	Token       string
}

var dependencies = []dependenciesModel{
	{
		Constructor: domain.GetRepository,
		Token:       "EXAMPLE-REPOSITORY",
	},
	{
		Constructor: domain.GetService,
		Token:       "EXAMPLE-SERVICE",
	},
	{
		Constructor: application.NewExampleApp,
		Token:       "EXAMPLE-APP",
	},
	{
		Constructor: interfaces.NewExampleHandler,
		Token:       "EXAMPLE-HANDLER",
	},
}

// Module binds all modules to the container
func Module(container *dig.Container) *server.Module {
	for _, v := range dependencies {
		if err := container.Provide(v.Constructor, dig.Name(v.Token)); err != nil {

		}
	}
	return &server.Module{
		Group:  "example",
		Routes: interfaces.ExampleRoutes(container),
	}
}
