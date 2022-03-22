// https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package http_server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) { // The Handler interface is what we need to implement in order to make a server, and the HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. HandlerFunc documentation: https://pkg.go.dev/net/http#HandlerFunc and Handler documentation: https://pkg.go.dev/net/http#Handler
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.Store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score) // ResponseWriter also implements io Writer so we can use fmt.Fprint to send strings as HTTP responses.
}

func GetPlayerScore(name string) int {
	switch name {
	case "Pepper":
		return 20
	case "Floyd":
		return 10
	}
	return 0
}
