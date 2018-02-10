package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/grpc/status"
)

func Run(hostBind string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	log.Println("Listening on " + hostBind)
	if err := http.ListenAndServe(hostBind, mux); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	msg, err := handleForward(context.Background(), r)
	if err != nil {
		status, _ := status.FromError(err)
		DefaultErrorHandler(w, status)
	} else {
		marshaler := jsonpb.Marshaler{}
		if err := marshaler.Marshal(w, msg); err != nil {
			fmt.Fprintf(w, "%s", err)
		}
	}
}
