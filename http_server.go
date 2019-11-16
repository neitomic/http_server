package http_server

import (
	"fmt"
	"github.com/go-akka/configuration"
	"github.com/labstack/echo"
	echomware "github.com/labstack/echo/middleware"
	"github.com/thanhtien522/http_server/controller"
	"github.com/thanhtien522/http_server/middleware"
	"github.com/thanhtien522/http_server/response/writer"
	"strings"
)

type httpServer struct {
	config     *configuration.Config
	echoServer *echo.Echo
	respWriter writer.Writer
}

func NewHttpServer(conf *configuration.Config, respWriter writer.Writer) *httpServer {
	e := echo.New()
	if respWriter == nil {
		respWriter = &writer.JsonWriter{}
	}
	return &httpServer{
		config:     conf,
		echoServer: e,
		respWriter: respWriter,
	}
}

func (s *httpServer) WithCORS(origins ...string) *httpServer {
	if len(origins) <= 0 {
		origins = []string{"*"}
	}
	cors := middleware.CORSWithConfig(echomware.CORSConfig{
		AllowOrigins:     origins,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true,
	})
	return s.WithMiddleware(cors)
}

func (s *httpServer) WithDefaultLogger() *httpServer {
	return s.WithLogger(nil)
}
func (s *httpServer) WithLogger(conf *echomware.LoggerConfig) *httpServer {
	if conf == nil {
		conf = &echomware.DefaultLoggerConfig
	}
	return s.WithMiddleware(echomware.LoggerWithConfig(*conf))
}

func (s *httpServer) WithErrorHandler(handler echo.HTTPErrorHandler) *httpServer {
	s.echoServer.HTTPErrorHandler = handler
	return s
}

func (s *httpServer) WithDefaultErrorHandler() *httpServer {
	return s.WithErrorHandler(controller.DefaultHttpErrorHandler(s.respWriter))
}

func (s *httpServer) WithMiddleware(middlewareFunc echo.MiddlewareFunc) *httpServer {
	s.echoServer.Use(middlewareFunc)
	return s
}

func (s *httpServer) RegisterRenderer(renderer echo.Renderer) *httpServer {
	s.echoServer.Renderer = renderer
	return s
}

func (s *httpServer) RegisterControllerGroup(prefix string, ctl controller.HttpController) *httpServer {
	g := s.echoServer.Group(prefix)
	for _, handler := range ctl.GetHandlers() {
		g.Add(
			strings.ToUpper(handler.Method),
			handler.Path,
			s.handlerFuncWrapper(handler),
			handler.MiddlewareFunc...
		)
	}
	return s
}

func (s *httpServer) RegisterController(ctl controller.HttpController) *httpServer {
	for _, handler := range ctl.GetHandlers() {
		s.echoServer.Add(
			strings.ToUpper(handler.Method),
			handler.Path,
			s.handlerFuncWrapper(handler),
			handler.MiddlewareFunc...
		)
	}
	return s
}

func (s *httpServer) WithStatic(path string, rootDir string) *httpServer {
	s.echoServer.Static(path, rootDir)
	return s
}

func (s *httpServer) Start() {
	bindHost := s.config.GetString("http_server.host", "0.0.0.0")
	bindPort := s.config.GetInt32("http_server.port", 8080)
	s.echoServer.Logger.Fatal(s.echoServer.Start(fmt.Sprintf("%s:%d", bindHost, bindPort)))
}

func (s *httpServer) handlerFuncWrapper(handler *controller.Handler) echo.HandlerFunc {
	respWriter := s.respWriter
	if handler.RespWriter != nil {
		respWriter = handler.RespWriter
	}
	return func(ctx echo.Context) error {
		if data, err := handler.HandleFunc(ctx); err != nil {
			return err
		} else {
			return respWriter.Write(ctx, data)
		}
	}
}