package main

import (
	"context"
	"encoding/json"
)

// Service will provide the interface so we can call to these methods
type Service interface {
	Status(ctx context.Context) (string, error)
	postInfo(ctx context.Context, message json.RawMessage) (bool, error)
}

// baseService will be used to group methods, so we can hide their implementation
type baseService struct{}

// ServiceConstructor creates a new instance masking the internal logic
func ServiceConstructor() Service {
	return baseService{}
}

// Status is the function that will return just "ok", it can be used for monitoring
func (baseService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

// postInfo is the function that will receive a json raw message and for now just print
// TODO: Implement logic that will enable to parse this message to an AMQP and return true if everything is okay
func (baseService) postInfo(ctx context.Context, message json.RawMessage) (bool, error) {
	//We can implement validations if message is empty here, but for now we dont need them
	amqpService := baseAQMPService{}
	published := amqpService.publishMessage(message)
	if published != false {
		return true, nil
	}
	return false, nil
}
