package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Main method that will serve the http server created in the server.go file
// Using the flag declaration we can use this to set an specific port if we have the usual 8080
// already in use.
func main() {
	var (
		httpAddress = flag.String("http", ":10000", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	srv := ServiceConstructor()
	// Using channel we can create goroutines that will allow us to to send values and receive them
	channelError := make(chan error)

	// First goroutine used when the user wants to exit the server, we send a SIGTERM signal so the process
	// stops when the user presses CTRL C
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		channelError <- fmt.Errorf("%s", <-c)
	}()

	endpoints := Endpoints{
		StatusEndpoint:   buildStatusEndpoint(srv),
		postInfoEndpoint: buildpostInfoEndpoint(srv),
	}

	// This gorouting will be listening to every single incoming request received by the server
	// So if we call it via postman or CURL this goroutine will be fired
	// Since this goroutine is using the handler we created in the server.go file
	// it will implement the build methods in endpoints and respond with the correct objects
	// to every single request
	go func() {
		log.Println("Base Service is listening on port:", *httpAddress)
		handler := Server(ctx, endpoints)
		channelError <- http.ListenAndServe(*httpAddress, handler)
	}()

	log.Fatalln(<-channelError)
}
