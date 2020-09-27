package main

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// I decided to implement mux as the request router and dispatcher,
// I took this decision because I wanted to look for something that would allow me to
// create methods with specific decoders and encoders in case we need to switch them later on.

// Server function will be used to create an http server using the endpoints declared
func Server(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware) // Using this we can set globally all the endpoints with Content-Type JSON
	r.Methods("GET").Path("/status").Handler(httptransport.NewServer(
		endpoints.StatusEndpoint,
		decodeStatusRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/postInfo").Handler(httptransport.NewServer(
		endpoints.postInfoEndpoint,
		decodepostInfoRequest,
		encodeResponse,
	))
	return r
}

// commonMiddleWare is used to expand the Router functionality of mux
// in this case I am using it to add to every single response of the endpoints
// the content-type json, so others developers dont have to do it manually
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
