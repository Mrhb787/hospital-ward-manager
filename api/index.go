package api

import (
	_ "embed"
	"net/http"

	"github.com/Mrhb787/hospital-ward-manager/service/http/health"
	"github.com/Mrhb787/hospital-ward-manager/transport"
)

//go:embed main.html
var mainHTML []byte

func Health(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		// Serve index.html
		serveMainHTML(w)
		return
	}

	// initalize services
	healthService := health.NewService()

	// http handler
	h := transport.NewHandler(healthService)

	// serve http
	h.ServeHTTP(w, r)

}

func serveMainHTML(w http.ResponseWriter) {
	// Set Content-Type header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Write the embedded HTML to the response
	w.Write(mainHTML)
}
