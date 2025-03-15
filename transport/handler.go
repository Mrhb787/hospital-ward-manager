package transport

import (
	"net/http"

	"github.com/Mrhb787/hospital-ward-manager/common"
	"github.com/Mrhb787/hospital-ward-manager/service/http/database"
	"github.com/Mrhb787/hospital-ward-manager/service/http/health"
)

type CustomHttpHandler interface {
	NewHandler() http.Handler
}

type HttpHandlerRequest struct {
	HealthService health.Service
	DbService     database.Service
	DbClient      database.Client
}

type vercelHttpHandler struct {
	healthService health.Service
	DbClient      database.Client
}

type baseHttpHandler struct {
	healthService health.Service
	dbService     database.Service
}

func GetHttpHandler(flag common.HttpHandlerType, req HttpHandlerRequest) CustomHttpHandler {
	switch flag {
	case common.VERCEL:
		return &vercelHttpHandler{healthService: req.HealthService, DbClient: req.DbClient}
	case common.BASE:
		return &baseHttpHandler{healthService: req.HealthService, dbService: req.DbService}
	}
	return nil
}

func (h *baseHttpHandler) NewHandler() http.Handler {
	// router
	r := common.NewRouter()

	// endpoints initialize
	healthEndpoints := MakeHealthEndpoints(h.healthService)

	// health check
	r.Methods(common.GET.ToString()).Path("/health").Handler(common.NewServer(
		healthEndpoints.HeathCheckEndpoint,
		DecodeHealthRequest,
		EncodeHealthResponse,
		optionsWithoutRouteCheck...,
	))

	return r
}

func (h *vercelHttpHandler) NewHandler() http.Handler {
	// router
	r := common.NewRouter()

	// endpoints initialize
	healthEndpoints := MakeHealthEndpoints(h.healthService)

	// health check
	r.Methods(common.GET.ToString()).Path("/health").Handler(common.NewServer(
		healthEndpoints.HeathCheckEndpoint,
		DecodeHealthRequest,
		EncodeHealthResponse,
		optionsWithoutRouteCheck...,
	))

	return r
}
