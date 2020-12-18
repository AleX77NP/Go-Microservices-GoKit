package todo

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	//AddTodoRequest ...
	AddTodoRequest struct {
		ID string `json:"id"`
		Text string `json:"text"`
		Completed bool `json:"completed"`
	}
	//AddTodoResponse ...
	AddTodoResponse struct {
		Ok string `json:"ok"`
	}
	// GetByIDRequest ...
	GetByIDRequest struct {
		ID string `json:"id"`
	}
	// GetByIDResponse ...
	GetByIDResponse struct {
		Text string `json:"text"`
	}
	//DeleteTodoRequest ...
	DeleteTodoRequest struct {
		ID string `json:"id"`
	}
	//DeleteTodoResponse ...
	DeleteTodoResponse struct {
		Deleted string `json:"deleted"`
	}
	//UpdateTodoRequest ...
	UpdateTodoRequest struct {
		ID string `json:"id"`
		Text string `json:"text"`
		Completed bool `json:"completed"`
	}
	//UpdateTodoResponse ...
	UpdateTodoResponse struct {
		Updated string `json:"updated"`
	}
)

func encodeReponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeTodo(ctx context.Context, r *http.Request) (interface{}, error) {
	var req AddTodoRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}