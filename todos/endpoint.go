package todo 

import (
	"github.com/go-kit/kit/endpoint"
	"context"
)

//Endpoints ...
type Endpoints struct {
	Add endpoint.Endpoint
	GetByID endpoint.Endpoint
	Delete endpoint.Endpoint
	Update endpoint.Endpoint
}

// MakeEndpoints ...
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Add: makeAddEndpoint(s),
		GetByID: makeGetByIDEndpoint(s),
		Delete: makeDeleteEndPoint(s),
		Update: makeUpdateEndpoint(s),
	}
}

func makeAddEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddTodoRequest)
		ok, err := s.Add(ctx, req.ID, req.Text, req.Completed)
		return AddTodoResponse{Ok: ok}, err
	}
}

func makeGetByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)
		text, err := s.GetByID(ctx,req.ID)
		return GetByIDResponse{
			Text: text,
		}, err
	}
}

func makeDeleteEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddTodoRequest)
		err := s.Delete(ctx, req.ID)
		return DeleteTodoResponse{Deleted: "deleted"},err
	}
}

func makeUpdateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddTodoRequest)
		_, err := s.Update(ctx, req.ID, req.Text, req.Completed)
		return UpdateTodoResponse{Updated:  "updated"},err
	}
}