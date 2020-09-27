package main

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints will be used to expose the methods written in the transport layer and service layer
type Endpoints struct {
	StatusEndpoint   endpoint.Endpoint
	postInfoEndpoint endpoint.Endpoint
}

// These build methods are used to call the correct service and respond with a proper object

// buildStatusEndpoint will use the status request and status response declared in the transport layer
// which consumes the logic declared in the service file
func buildStatusEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(statusRequest)
		s, err := srv.Status(ctx)
		if err != nil {
			return statusResponse{s}, err
		}
		return statusResponse{s}, nil
	}
}

// buildpostInfoEndpoint will use the postinfo getter and setter declared in the transport layer
func buildpostInfoEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postInfoRequest)
		b, err := srv.postInfo(ctx, req.Message)
		if err != nil {
			return postInfoResponse{b, err.Error()}, nil
		}
		return postInfoResponse{b, ""}, nil
	}
}

// Status endpoint mapping
func (e Endpoints) Status(ctx context.Context) (string, error) {
	req := statusRequest{}
	resp, err := e.StatusEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	statusResp := resp.(statusResponse)
	return statusResp.Status, nil
}

// postInfo endpoint mapping
func (e Endpoints) postInfo(ctx context.Context, message json.RawMessage) (bool, error) {
	req := postInfoRequest{Message: message}
	resp, err := e.postInfoEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	postInfoResp := resp.(postInfoResponse)
	if postInfoResp.Err != "" {
		return false, errors.New(postInfoResp.Err)
	}
	return postInfoResp.Valid, nil
}
