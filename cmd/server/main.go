package main

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/di"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/config"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/dig"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/http/server"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/jwt"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/log"
)

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
	var logg *log.Logger
	container := dig.BuildContainer()

	if err = di.ConfiDI(container); err != nil {
		logg.Fatal("error in modules", zap.Error(err))
	}

	if err = container.Invoke(func(logger *log.Logger) {
		logg = logger
	}); err != nil {
		fmt.Println("error in logger", zap.Error(err))
	}

	if err = container.Invoke(func(httpServer *server.Server, logger *log.Logger, jwt *jwt.JWT, conf *config.Config) {
		if err := dig.ResgisterModules(container, httpServer.Router, logger); err != nil {
			logger.Fatal("error in resgister modules", zap.Error(err))
		}

		if err = httpServer.Run(logger, jwt, conf); err != nil {
			logger.Fatal("error in init server", zap.Error(err))
		}
	}); err != nil {
		logg.Fatal("error in server", zap.Error(err))
	}
}
