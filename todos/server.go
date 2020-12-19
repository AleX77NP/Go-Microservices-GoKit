package todo

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

// NewHTTPServer ...
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleWare)

	r.Methods("POST").Path("/todo").Handler(httptransport.NewServer(
		endpoints.Add,
		decodeTodo,
		encodeReponse,
	))

	r.Methods("GET").Path("/todo/{id}").Handler(httptransport.NewServer(
		endpoints.GetByID,
		decodeIDReq,
		encodeReponse,
	))

	r.Methods("DELETE").Path("/todo/{id}").Handler(httptransport.NewServer(
		endpoints.Delete,
		decodeDeleteReq,
		encodeReponse,
	))

	r.Methods("PUT").Path("/todo/{id}").Handler(httptransport.NewServer(
		endpoints.Update,
		decodeUpdateReq,
		encodeReponse,
	))

	return r
}

func commonMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w,r)
	})
}