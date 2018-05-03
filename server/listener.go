package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/api-gateway/loader"
	"github.com/api-gateway/types"
	"github.com/api-gateway/types/log"
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
	mux.HandleFunc("/loader", loaderHandler)

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
		w.Write(msg)
	}
}

func loaderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		out := ""
		for k, v := range loader.RuleStore {
			out = out + "\n" + k + " --> " + v.Package + "." + v.Service + "." + *v.Method.Name
		}
		w.Write([]byte(out))
	}

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error(err)
		}
		var methods []types.MethodWrapper
		err = json.Unmarshal(body, &methods)
		if err != nil {
			log.Error(err)
		}

		for _, md := range methods {
			key := md.Pattern.Verb + ":" + md.Pattern.Path
			loader.RuleStore[key] = md
		}
	}
}
