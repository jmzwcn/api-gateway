package main

import (
	"github.com/api-gateway/common"
	"github.com/api-gateway/loader"
	"github.com/api-gateway/server"
)

func main() {
	log.Debug("API Gateway Start...")
	loader.ParseAndLoad()
	server.Listen()
}
