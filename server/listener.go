package server

import (
	"context"
	"io"
	"log"
	"net/http"
)

func Listen() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	if err := http.ListenAndServe(":6060", mux); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := handleForward(ctx, r)
	if err != nil {
		log.Println(err)
		io.WriteString(w, err.Error())
	} else {
		io.WriteString(w, "Response: "+resp)
	}
}