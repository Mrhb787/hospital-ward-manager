package api

import (
	"net/http"

	"github.com/Mrhb787/hospital-ward-manager/common"
	"github.com/Mrhb787/hospital-ward-manager/service/http/health"
	"github.com/Mrhb787/hospital-ward-manager/transport"
)

func Health(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		// Serve index.html
		common.ServeIndexHTML(w, r)
		return
	}

	// initalize services
	healthService := health.NewService()

	// http handler
	h := transport.NewHandler(healthService)

	// serve http
	h.ServeHTTP(w, r)

}
