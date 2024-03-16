package httpserver

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// route sets the handler for the path "/"
func route(a api) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.routeHandlers)
	mux.Handle("/docs/", httpSwagger.WrapHandler)

	return mux
}
