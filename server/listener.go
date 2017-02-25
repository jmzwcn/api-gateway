package server

import (
	"api-gateway/common"
	"api-gateway/config"
	"context"
	"io"
	"net/http"
)

func Listen() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	port := config.NewConfiguration().Port
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := handleForward(ctx, r)
	if err != nil {
		log.Error(err)
		io.WriteString(w, err.Error())
	} else {
		io.WriteString(w, "Response: "+resp)
	}
}
