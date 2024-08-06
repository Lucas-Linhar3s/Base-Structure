package di

import (
	"go.uber.org/dig"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/modules/example"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/http/server"
)

func Modules(container *dig.Container) []server.Module {
	return []server.Module{
		*example.Module(container),
	}
}
