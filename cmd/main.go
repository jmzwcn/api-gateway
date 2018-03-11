package main

import (
	"flag"
	"os"

	"github.com/api-gateway/server"
	"github.com/api-gateway/types/log"
)

var (
	hostBind string
)

func main() {
	log.Println("API-Gateway start...")
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
