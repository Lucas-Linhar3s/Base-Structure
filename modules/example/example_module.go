package example

import (
	"fmt"

	"go.uber.org/dig"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/database"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/modules/example/application"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/modules/example/domain"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/modules/example/interfaces"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/http/server"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/log"
)

// Binds binds all modules to the container
func Binds(container *dig.Container) *server.Module {
	//TODO Implementar os modulos
	var logg *log.Logger
	// Registra os modulos
	if err := container.Invoke(func(logger *log.Logger) {
		logg = logger
	}); err != nil {
		logg.Error(err.Error())
		return nil
	}

	if err := container.Provide(func(db *database.Database) domain.IExample {
		return domain.GetRepository(db)
	}); err != nil {
		logg.Error(err.Error())
		return nil
	}

	if err := container.Provide(func(repo domain.IExample) *domain.Service {
		return domain.GetService(repo)
	}); err != nil {
		logg.Error(err.Error())
		return nil
	}

	if err := container.Provide(func(db *database.Database, service *domain.Service, logger *log.Logger) *application.ExampleApp {
		return application.NewExampleApp(db, service, logger)
	}); err != nil {
		logg.Error(err.Error())
		return nil
	}

	if err := container.Provide(func(app *application.ExampleApp) *interfaces.ExampleHandler {
		return interfaces.NewExampleHandler(app)
	}); err != nil {
		logg.Error(err.Error())
		return nil
	}

	// Registra as rotas
	return routes(container)
}

func routes(container *dig.Container) *server.Module {
	var h *interfaces.ExampleHandler
	if err := container.Invoke(func(handler *interfaces.ExampleHandler) {
		h = handler
	}); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	// Registra as rotas
	return &server.Module{
		Name: "example",
		Routes: []server.Route{
			{
				Method:      "GET",
				Path:        "/find",
				Handler:     h.Find,
				Description: "Find example",
				Middlewares: nil,
			},
		},
	}
}
