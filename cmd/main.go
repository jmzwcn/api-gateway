package main

import (
	"flag"
	"log"
	"os"

	"github.com/api-gateway/loader"
	"github.com/api-gateway/server"
)

var hostBind string

func main() {
	log.Println("API-Gateway start...")
	loader.Services()
	server.Run(hostBind)
}

func init() {
	var isHelp bool
	flag.StringVar(&hostBind, "bind", ":8080", "Bind address")
	flag.BoolVar(&isHelp, "help", false, "Print this help")
	flag.Parse()

	if isHelp {
		showHelp()
	}
}

func showHelp() {
	flag.PrintDefaults()
	os.Exit(1)
}
