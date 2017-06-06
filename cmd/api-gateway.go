package main

import (
	"flag"
	"os"

	"github.com/api-gateway/common"
	"github.com/api-gateway/loader"
	"github.com/api-gateway/server"

	_ "github.com/api-gateway/example/echo/api"
)

var hostBind, logLevel string

func main() {
	log.Info("API-Gateway start...")
	loader.ParseAndLoad()
	server.Listen(hostBind)
}

func init() {
	var isHelp bool
	flag.StringVar(&hostBind, "bind", ":8080", "Bind address")
	flag.StringVar(&logLevel, "logLevel", "debug", "Log Level")
	flag.BoolVar(&isHelp, "help", false, "Print this help")
	flag.Parse()

	log.SetLevel(logLevel)
	if isHelp {
		showHelp()
	}
}

func showHelp() {
	flag.PrintDefaults()
	os.Exit(1)
}
