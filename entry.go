package main

import (
	"api-gateway/common"
	"api-gateway/loader"
	"api-gateway/server"
)

func main() {
	log.Debug("API Gateway Start...")
	loader.ParseAndLoad()
	server.Listen()
}
