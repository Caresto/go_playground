package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// AMQPService will provide the interface so we can call to these methods
type AMQPService interface {
	publishMessage(body json.RawMessage) bool
}

// baseAQMPService will be used to group methods, so we can hide their implementation
type baseAQMPService struct{}

// AQMPServiceConstructor creates a new instance masking the internal logic
func AQMPServiceConstructor() AMQPService {
	return baseAQMPService{}
}

func (baseAQMPService) publishMessage(messageBody json.RawMessage) bool {
	if err := publish(messageBody); err != true {
		log.Fatalf("Error publishing in queue %v", err)
		return false
	}

	return true
}

// publish function will be the one Dialing to rabbitMQ instance and then publishing message
func publish(messageBody json.RawMessage) bool {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Cannot connect to brokers, please initialize the RabbitMQ instance first")
		return false
	}

	// After connection we open the channel to our MQ instance
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	// Now that the channel is open we can interact with the MQ Instance
	// and publish messages
	// QueueDeclare has 6 attributes in this order: Name, Durable, Exclusive, Auto-delete, other arguments
	q, err := ch.QueueDeclare(
		"MyQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	// Printing the queue for debugging purposes
	fmt.Println(q)
	// Handle any errors if we were unable to create the queue
	if err != nil {
		fmt.Println(err)
	}

	// Publishing our message received at the endpoint postInfo
	err = ch.Publish(
		"",
		"MyQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(messageBody),
		},
	)

	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println("Published the message to the queue")
	return true
}
