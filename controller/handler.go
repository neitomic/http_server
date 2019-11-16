package controller

import (
	"github.com/labstack/echo"
	"github.com/thanhtien522/http_server/response/writer"
)

type HandlerFunc func(echo.Context) (interface{}, error)

type Handler struct {
	Method         string
	Path           string
	HandleFunc     HandlerFunc
	MiddlewareFunc []echo.MiddlewareFunc
	RespWriter     writer.Writer
}

func NewHandler(method string,
	path string,
	handleFunc HandlerFunc,
	middlewares []echo.MiddlewareFunc,
	writer writer.Writer) *Handler {
	return &Handler{
		Method:         method,
		Path:           path,
		HandleFunc:     handleFunc,
		MiddlewareFunc: middlewares,
		RespWriter:     writer,
	}
}

type HttpController interface {
	GetHandlers() []*Handler
}