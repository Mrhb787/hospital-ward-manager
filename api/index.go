package api

import (
	"net/http"

	"github.com/Mrhb787/hospital-ward-manager/service/http/health"
	"github.com/Mrhb787/hospital-ward-manager/transport"
)

func Health(w http.ResponseWriter, r *http.Request) {

	// initalize services
	healthService := health.NewService()

	// http handler
	h := transport.NewHandler(healthService)

	// serve http
	h.ServeHTTP(w, r)

}
