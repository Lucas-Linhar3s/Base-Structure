package di

import (
	"flag"

	d "go.uber.org/dig"
	"go.uber.org/zap"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/database"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/modules/example"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/config"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/http/server"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/jwt"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/log"
)

// Modules register modules
func modules(container *d.Container) (err error) {
	// Registra os modulos
	if err := container.Provide(func(jwt *jwt.JWT, logger *log.Logger) (modules []server.Module) {
		modules = []server.Module{
			*example.Binds(container),
		}
		return
	}); err != nil {
		return err
	}

	return err
}

func ConfiDI(container *d.Container) error {
	if err := registerConfig(container); err != nil {
		return err
	}

	if err := registerJWT(container); err != nil {
		return err
	}

	if err := registerLogger(container); err != nil {
		return err
	}

	if err := registerServer(container); err != nil {
		return err
	}

	if err := registerDatabase(container); err != nil {
		return err
	}

	if err := modules(container); err != nil {
		return err
	}

	return nil
}

func registerConfig(container *d.Container) error {
	return container.Provide(func() *config.Config {
		var envConf = flag.String("conf", "../../config/local.yml", "config path, eg: -conf ../../config/local.yml")
		flag.Parse()
		confViper := config.NewViper(*envConf)
		return config.LoadAttributes(confViper)
	})
}

func registerJWT(container *d.Container) error {
	return container.Provide(func(conf *config.Config) *jwt.JWT {
		return jwt.NewJwt(conf)
	})
}

func registerLogger(container *d.Container) error {
	return container.Provide(func(conf *config.Config) *log.Logger {
		return log.NewLog(conf)
	})
}

func registerServer(container *d.Container) error {
	return container.Provide(func(jwt *jwt.JWT, logger *log.Logger, conf *config.Config) *server.Server {
		return server.NewServer(container)
	})
}
func registerDatabase(container *d.Container) error {
	return container.Provide(func(conf *config.Config, logger *log.Logger) *database.Database {
		db, err := database.Open(conf, "", "sqlite")
		if err != nil {
			logger.Fatal("failed to open database", zap.Error(err))
		}
		return db
	})
}
