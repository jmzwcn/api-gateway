package server

import (
	"context"
	"net/http"

	"google.golang.org/grpc/status"

	"github.com/api-gateway/common"
)

func Listen(hostBind string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	log.Info("Listening on " + hostBind)
	if err := http.ListenAndServe(hostBind, mux); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := handleForward(ctx, r)
	if err != nil {
		log.Error(err)
		status, _ := status.FromError(err)
		http.Error(w, err.Error(), HTTPStatusFromCode(status.Code()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(resp))
	}
}
