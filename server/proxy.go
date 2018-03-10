package server

import (
	"net/http"
	"strings"

	"golang.org/x/net/websocket"
)

type ExServeMux struct {
	http.ServeMux
}

func (mux *ExServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.ToLower(r.Header.Get("Upgrade")) != "websocket" {
		mux.ServeMux.ServeHTTP(w, r)
		return
	}

	wsh := websocket.Handler(func(ws *websocket.Conn) { streamHandler(r, ws) })
	wsh.ServeHTTP(w, r)
}
