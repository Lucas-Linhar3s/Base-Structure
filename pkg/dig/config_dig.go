package dig

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/http/server"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/log"
)

// BuildContainer builds the dependency injection container
func BuildContainer() *dig.Container {
	container := dig.New()
	return container
}

// ResgisterModules register modules
func ResgisterModules(container *dig.Container, router *gin.Engine, logger *log.Logger) error {
	return container.Invoke(func(modules []server.Module) {
		for _, module := range modules {
			module.Register(router, logger)
		}
	})
}

// InvokeService invoca um serviço do tipo especificado no container e retorna a instância correspondente
func InvokeService[T any](container *dig.Container, serviceType *T) (*T, error) {
	var result *T

	// Utilize container.Invoke para obter o serviço do tipo especificado
	err := container.Invoke(func(service T) {
		result = &service
	})

	if err != nil {
		return result, err
	}

	if result == nil {
		return result, errors.New("service not found")
	}

	return result, nil
}
