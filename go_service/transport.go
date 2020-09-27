package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// This part of the file will be used to map all requests and response
// All responses will be handled by JSON as its a standard nowadays
// statusRequest will be empty since it doesn't require anything
type statusRequest struct{}

// statusResponse will respond a json structured answer to the user when it calls the request
type statusResponse struct {
	Status string `json: "status"`
}

// postInfoRequest will receive a json structured message, this will be read from a key called Message
type postInfoRequest struct {
	Message json.RawMessage `json:"message"`
}

// postInfoResponse will respond a json structured answer to the user when it is valid, and it will add an error
// if something goes wrong
type postInfoResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err, omitempty`
}

// These decoders methods will be used for all of the incoming messages can be done via curl or postman
func decodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req statusRequest
	return req, nil
}

func decodepostInfoRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req postInfoRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// This function will be used to encode the response output, since we are responding only json we will use a
// JSOn encoder but we can expand this if its required
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
