package di

import (
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/database"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/config"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/http/server"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/jwt"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/log"
)

type DependenciesModel struct {
	Constructor interface{}
	Token       string
}

var Dependencies = []DependenciesModel{
	{
		Constructor: config.NewViper,
		Token:       "VIPER",
	},
	{
		Constructor: config.LoadAttributes,
		Token:       "CONFIG",
	},
	{
		Constructor: jwt.NewJwt,
		Token:       "JWT",
	},
	{
		Constructor: log.NewLog,
		Token:       "LOGGER",
	},
	{
		Constructor: server.NewServer,
		Token:       "SERVER",
	},
	{
		Constructor: database.NewDatabase,
		Token:       "DATABASE",
	},
}
