package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/modules/example/application"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/http/responses"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/utils"
)

// ExampleHandler represents the handler for the example module
type ExampleHandler struct {
	app *application.ExampleApp
}

// NewExampleHandler returns a new ExampleHandler
func NewExampleHandler(app *application.ExampleApp) *ExampleHandler {
	return &ExampleHandler{
		app: app,
	}
}

// @Summary Find example
// @Description Find example by name
// @Tags example
// @Produce json
// @Param name query string true "Name"
// @Success 200 {object} application.ExampleRes
// @Router /example/find [get]
func (h *ExampleHandler) Find(ctx *gin.Context) {
	var req application.ExampleReq
	req.Name = utils.GetStringPointer(ctx.Query("name"))
	res, err := h.app.Find(ctx, req)
	if err != nil {
		responses.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	responses.HandleSuccess(ctx, res)
}
