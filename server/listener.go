package server

import (
	"context"
	"io"
	"net/http"

	"google.golang.org/grpc/status"

	"github.com/api-gateway/common"
	"github.com/api-gateway/config"
)

func Listen(hostBind string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	if hostBind == "" {
		hostBind = ":" + config.NewConfiguration().Port
	}

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
		log.Debug(status)
		w.WriteHeader(HTTPStatusFromCode(status.Code()))
		//w.WriteHeader(500)
		io.WriteString(w, err.Error())
	} else {
		//w.WriteHeader(http.StatusOK)
		io.WriteString(w, resp)
	}
}
