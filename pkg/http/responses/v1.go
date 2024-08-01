package responses

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleSuccess(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	resp := map[string]interface{}{"Code": errorCodeMap[ErrSuccess], "Message": ErrSuccess.Error(), "Data": data}
	if _, ok := errorCodeMap[ErrSuccess]; !ok {
		resp = map[string]interface{}{"Code": 0, "Message": "", "Data": data}
	}
	ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx *gin.Context, httpCode int, err error, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := map[string]interface{}{"Code": httpCode, "Message": err.Error(), "Data": data}
	if _, ok := errorCodeMap[ErrSuccess]; !ok {
		resp = map[string]interface{}{"Code": 500, "Message": "unknown error", "Data": data}
	}
	ctx.JSON(httpCode, resp)
}

type Error struct {
	Code    int
	Message string
}

var errorCodeMap = map[error]int{}

func newError(code int, msg string) error {
	err := errors.New(msg)
	errorCodeMap[err] = code
	return err
}
func (e Error) Error() string {
	return e.Message
}
