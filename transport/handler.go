package transport

import (
	"net/http"

	"github.com/Mrhb787/hospital-ward-manager/common"
	"github.com/Mrhb787/hospital-ward-manager/service/http/database"
	"github.com/Mrhb787/hospital-ward-manager/service/http/health"
)

type HttpHandlerRequest struct {
	HealthService health.Service
	DbService     database.Service
}

func NewHandler(req HttpHandlerRequest) http.Handler {
	// router
	r := common.NewRouter()

	// endpoints initialize

	// health endpoints
	healthEndpoints := MakeHealthEndpoints(req.HealthService)

	// service endpoints
	serviceEndpoints := MakeHttpServiceEndpoints(req.DbService)

	// health check
	r.Methods(common.GET.ToString()).Path("/health").Handler(common.NewServer(
		healthEndpoints.HeathCheckEndpoint,
		DecodeHealthRequest,
		EncodeHealthResponse,
		optionsWithoutRouteCheck...,
	))

	// get user by id
	r.Methods(common.GET.ToString()).Path("/user").Handler(common.NewServer(
		serviceEndpoints.GetUserEndpoint,
		DecodeGetUserByIdRequest,
		EncodeGenericResponse,
		optionsWithoutRouteCheck...,
	))

	return r
}
