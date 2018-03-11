package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/api-gateway/types"
	"github.com/api-gateway/types/log"
	"github.com/gogo/protobuf/jsonpb"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func init() {
	grpc.EnableTracing = true
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
}

func Run(hostBind string) {
	mux := new(ExServeMux)
	mux.HandleFunc("/", unaryHandler)
	mux.HandleFunc("/debug/requests", trace.Traces)
	mux.HandleFunc("/debug/events", trace.Events)

	log.Println("Listening on " + hostBind)
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
