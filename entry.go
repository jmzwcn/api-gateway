package main

import (
	"api-gateway/loader"
	"api-gateway/server"
	"log"
)

func main() {
	log.Println("API Gateway Begin...")
	loader.ParseAndLoad()
	server.Listen()
}
