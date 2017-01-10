package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func Listen() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GenericHandler)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func GenericHandler(w http.ResponseWriter, r *http.Request) {
	//will forward to gRPC
	io.WriteString(w, "hello, world!"+time.Now().String())
}
