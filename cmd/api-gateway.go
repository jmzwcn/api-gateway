package main

import (
	"api-gateway/loader"
	"api-gateway/server"

	"log"
)

func main() {
	log.Println("API Gateway Start...")
	loader.ParseAndLoad()
	server.Listen()
}
