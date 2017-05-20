package main

import (
	"github.com/api-gateway/common"
	"github.com/api-gateway/loader"
	"github.com/api-gateway/server"
)

func main() {
	log.Info("API-Gateway start...")
	loader.ParseAndLoad()
	server.Listen()
}
