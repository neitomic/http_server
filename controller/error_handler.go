package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/thanhtien522/http_server/response"
	"github.com/thanhtien522/http_server/response/writer"
	"net/http"
)

func DefaultHttpErrorHandler(writer writer.Writer) echo.HTTPErrorHandler {

	return func(err error, ctx echo.Context) {
		var (
			code = http.StatusInternalServerError
			msg  string
		)

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			msg = fmt.Sprintf("%v", he.Message)
			if he.Internal != nil {
				msg = fmt.Sprintf("%s, cause by: %v", msg, he.Internal)
			}
		} else {
			msg = err.Error()
		}

		ctx.Logger().Error(err)

		// Send response
		if !ctx.Response().Committed {
			if ctx.Request().Method == echo.HEAD { // Issue #608
				err = ctx.NoContent(code)
			} else {
				resp := response.ErrorResp(code, &response.APIError{Code: code, Message: msg})
				err = writer.Write(ctx, resp)
			}
			if err != nil {
				ctx.Logger().Error(err)
			}
		}
	}
}
