package transport

import (
	"context"

	"github.com/Mrhb787/hospital-ward-manager/common"
	"github.com/Mrhb787/hospital-ward-manager/service/http/health"
)

type Endpoints struct {
	HeathCheckEndpoint common.Endpoint
}

func MakeHealthEndpoints(s health.Service) Endpoints {
	return Endpoints{
		HeathCheckEndpoint: MakeHealthEndpoint(s),
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
