package writer

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/thanhtien522/http_server/response"
	"net/http"
)

type BinaryWriter struct{}

func (w *BinaryWriter) Write(ctx echo.Context, resp interface{}) error {
	switch r := resp.(type) {
	case response.BinaryResp:
		ctx.Response().Header().Add("Content-Type", r.ContentType)
		for key, values := range r.HttpHeader {
			for _, val := range values {
				ctx.Response().Header().Add(key, val)
			}
		}
		http.ServeContent(ctx.Response(), ctx.Request(), r.Name, r.ModTime, r.Reader)
		return nil
	default:
		return errors.New("unsupported response data type")
	}
}
