package application

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/database"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/modules/example/domain"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/log"
)

// ExampleApp represents the example application
type ExampleApp struct {
	db      *database.Database
	service *domain.Service
	logger  *log.Logger
}

// ExampleReq represents the example request
func NewExampleApp(db *database.Database, service *domain.Service, logger *log.Logger) *ExampleApp {
	return &ExampleApp{
		db:      db,
		service: service,
		logger:  logger,
	}
}

// Find example
func (a *ExampleApp) Find(ctx *gin.Context, req ExampleReq) (res ExampleRes, err error) {
	const msg = "failed to find example"
	var (
		data domain.ExampleModel
	)

	if err = copier.Copy(&data, req); err != nil {
		a.logger.Error(msg, zap.Error(err))
		return
	}

	res.Data, err = a.service.Find(data)
	if err != nil {
		a.logger.Error(msg, zap.Error(err))
		return
	}

	return
}
