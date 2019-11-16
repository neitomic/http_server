package ctl

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/thanhtien522/http_server/controller"
	"github.com/thanhtien522/http_server/response"
	"time"
)

type pingController struct{}

func NewPingController() *pingController {
	return &pingController{}
}

func (ctl *pingController) Ping(ctx echo.Context) (interface{}, error) {
	currentTime := time.Now()
	msg := fmt.Sprint("pong: ", currentTime.Format("2006-01-02 15:04:05.000000"))
	return response.SuccessResp(msg), nil
}

func (ctl *pingController) GetHandlers() []*controller.Handler {
	return []*controller.Handler{
		controller.NewHandler(echo.GET, "/ping", ctl.Ping, nil, nil),
	}
}
