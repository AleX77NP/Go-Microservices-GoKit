package todo 

import (
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Add endpoint.Endpoint
	GetByID endpoint.Endpoint
	Delete endpoint.Endpoint
	Update endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Add: makeAddEndpoint(s),
		GetByID: makeGetByIDEndpoint(s),
		Delete: makeDeleteEndPoint(s),
		Update: makeUpdateEndpoint(s),
	}
}