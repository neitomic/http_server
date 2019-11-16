package writer

import (
	"github.com/labstack/echo"
	"github.com/thanhtien522/http_server/response"
	"net/http"
)

type JsonWriter struct{}

func (w *JsonWriter) Write(ctx echo.Context, resp interface{}) error {
	switch r := resp.(type) {
	case response.HttpResp:
		for key, values := range r.HttpHeader {
			for _, val := range values {
				ctx.Response().Header().Add(key, val)
			}
		}
		return ctx.JSON(r.HttpCode, resp)
	default:
		return ctx.JSON(http.StatusOK, resp)
	}
}
