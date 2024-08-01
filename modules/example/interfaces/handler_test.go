package interfaces_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/di"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/middleware"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/dig"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/jwt"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/log"
)

func TestFind(t *testing.T) {
	container := dig.BuildContainer()
	if err := di.ConfiDI(container); err != nil {
		t.Fatal(err)
	}

	r := gin.Default()
	if err := container.Invoke(func(jwt *jwt.JWT, logger *log.Logger) {
		r.Use(
			middleware.RequestLogMiddleware(logger),
			middleware.CORSMiddleware(),
			middleware.NoStrictAuth(jwt, logger),
			middleware.ResponseLogMiddleware(logger),
		)
	}); err != nil {
		t.Fatal(err)
	}
	t.Log("TestFind")

	// group := r.Group("/example")
	// example.Binds()
	t.Run("TestFind", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/example/find", nil)

		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
