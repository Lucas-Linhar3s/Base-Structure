package server

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/config"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/jwt"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/log"
)

// Server http server
type Server struct {
	Router    *gin.Engine
	Container *dig.Container
}

// NewServer new server with router
func NewServer(container *dig.Container) *Server {
	return &Server{Router: gin.Default(), Container: container}
}

// Run server
func (s *Server) Run(logger *log.Logger,
	jwt *jwt.JWT, conf *config.Config) error {
	// swagger doc
	// docs.SwaggerInfo.BasePath = "/"
	// docs
	swag := s.Router.Group("/swagger")
	swag.GET("/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		// ginSwagger.DefaultModelsExpandDepth(-1),
		// ginSwagger.PersistAuthorization(true),
	))

	s.Router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found " + " : " + c.Request.URL.String()})
	})
	return s.Router.Run(conf.Http.Port)
}
