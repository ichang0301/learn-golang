// https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package http_server

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) { // Handler documentation: https://pkg.go.dev/net/http#Handler
	switch r.URL.Path {
	case "/players/Pepper":
		fmt.Fprint(w, "20") // ResponseWriter also implements io Writer so we can use fmt.Fprint to send strings as HTTP responses.
	case "/players/Floyd":
		fmt.Fprint(w, "10") // ResponseWriter also implements io Writer so we can use fmt.Fprint to send strings as HTTP responses.
	}
}
