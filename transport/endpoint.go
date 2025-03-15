package transport

import (
	"context"

	"github.com/Mrhb787/hospital-ward-manager/common"
	"github.com/Mrhb787/hospital-ward-manager/service/http/database"
	"github.com/Mrhb787/hospital-ward-manager/service/http/health"
)

type Endpoints struct {
	HeathCheckEndpoint common.Endpoint
	GetUserEndpoint    common.Endpoint
}

func MakeHealthEndpoints(s health.Service) Endpoints {
	return Endpoints{
		HeathCheckEndpoint: MakeHealthEndpoint(s),
	}
}

func MakeHttpServiceEndpoints(dbService database.Service) Endpoints {
	return Endpoints{
		GetUserEndpoint: MakeGetUserEndpoint(dbService),
	}
}

func MakeHealthEndpoint(s health.Service) common.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status := s.Health(ctx)
		return healthResponse{
			Status: status,
		}, err
	}
}

func MakeGetUserEndpoint(dbService database.Service) common.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		resp, err := dbService.GetUserById(1)
		return resp, err
	}
}
