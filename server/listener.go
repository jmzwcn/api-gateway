package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/api-gateway/types"
	"github.com/api-gateway/types/log"
	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/grpc/status"
)

func Run(hostBind string) {
	mux := new(ExServeMux)
	mux.HandleFunc("/", unaryHandler)

	log.Infoln("Listening on " + hostBind)
	if err := http.ListenAndServe(hostBind, mux); err != nil {
		log.Fatal(err)
	}
}

func unaryHandler(w http.ResponseWriter, r *http.Request) {
	msg, err := handleForward(context.Background(), r)
	if err != nil {
		status, _ := status.FromError(err)
		types.DefaultErrorHandler(w, status)
	} else {
		marshaler := jsonpb.Marshaler{}
		if err := marshaler.Marshal(w, msg); err != nil {
			fmt.Fprintf(w, "%s", err)
		}
	}
}
