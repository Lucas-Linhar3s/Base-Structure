package main

import (
	d "go.uber.org/dig"
	"go.uber.org/zap"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/di"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/config"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/dig"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/http/server"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/jwt"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/log"
)

type dependencies struct {
	d.In
	Log        *log.Logger    `name:"LOGGER"`
	HTTPServer *server.Server `name:"SERVER"`
	JWT        *jwt.JWT       `name:"JWT"`
	Config     *config.Config `name:"CONFIG"`
}

// @title           Modularize example API
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8000
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	var err error
	var dep *dependencies
	container := dig.BuildContainer()

	if err = di.RegisterDI(container); err != nil {
		panic(err)
	}

	if dep, err = dig.InvokeService(container, dep); err != nil {
		panic(err)
	}

	if err := dig.ResgisterModules(container, dep.HTTPServer.Router, dep.Log); err != nil {
		dep.Log.Fatal("error in resgister modules", zap.Error(err))
	}

	if err = dep.HTTPServer.Run(dep.Log, dep.JWT, dep.Config); err != nil {
		dep.Log.Fatal("error in init server", zap.Error(err))
	}
}
