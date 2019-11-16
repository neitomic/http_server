package main

import (
	"github.com/go-akka/configuration"
	"github.com/thanhtien522/http_server"
	"github.com/thanhtien522/http_server/example/ctl"
	"github.com/thanhtien522/http_server/response/writer"
)

func main() {
	conf := configuration.LoadConfig("application.conf")

	pingCtl := ctl.NewPingController()
	server := http_server.NewHttpServer(conf, &writer.JsonWriter{})
	server.
		WithDefaultLogger().
		RegisterController(pingCtl).
		Start()
}
