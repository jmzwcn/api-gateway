package server

import (
	"context"
	"io"
	"net/http"

	"github.com/api-gateway/common"
	"github.com/api-gateway/config"
)

func Listen() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	port := config.NewConfiguration().Port
	log.Info("Listening on port:" + port)
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
		io.WriteString(w, resp)
	}
}
