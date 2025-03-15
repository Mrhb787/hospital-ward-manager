package transport

import (
	"net/http"

	"github.com/Mrhb787/hospital-ward-manager/common"
	"github.com/Mrhb787/hospital-ward-manager/service/http/health"
)

func NewHandler(healthService health.Service) http.Handler {
	// router
	r := common.NewRouter()

	// endpoints initialize
	healthEndpoints := MakeHealthEndpoints(healthService)

	// health check
	r.Methods(common.GET.ToString()).Path("/health").Handler(common.NewServer(
		healthEndpoints.HeathCheckEndpoint,
		DecodeHealthRequest,
		EncodeHealthResponse,
		optionsWithoutRouteCheck...,
	))

	return r
}
