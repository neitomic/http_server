package writer

import "github.com/labstack/echo"

type WriteFunc func(ctx echo.Context, data interface{}) error
type Writer interface {
	Write(ctx echo.Context, HttpResp interface{}) error
}
