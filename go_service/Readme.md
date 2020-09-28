Simple Go Service using Go-kit

Go-kit basic architecture goes like this

- Transport
- Operational metrics
- Balancing and limiting
- Business Analytics
- Application Logging
- Service metrics
- Business Logic

The service has two basic endpoints:
- GET Endpoint `http://localhost:10000/status` which is going to return the status of the service (It can be used for monitoring)
- POST Endpoint `http://localhost:10000/postInfo` which receives a json message as input and this message is sent to a rabbitMQ queue using AMQP(https://github.com/streadway/amqp)

In order to run this you will need to install the following go libraries using go cli:

```
go get -u github.com/streadway/amqp
go get -u github.com/go-logfmt/logfmt
go get -u github.com/gorilla/mux
go get -u github.com/go-kit/kit
```

In order to execute this example you will need to install Docker first

If you already have docker in your machine you are able continue

Install RabbitMQ via Docker for easy testing
```
docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

Clone the repository and after installing the dependencies the example can be executed with 
```
go run *.go
```

Running examples:
Posting a message to endpoint `http://localhost:10000/postInfo`
Terminal displaying service running and receiving a payload
https://www.dropbox.com/s/9gzwegjaec5u1y3/Screen%20Shot%202020-09-27%20at%205.42.34%20PM.png?dl=0

Message published in queue in the RabbitMQ interface `http://localhost:15672/#/queues/%2F/MyQueue`
https://www.dropbox.com/s/j3ev5kg4o82cs4r/Screen%20Shot%202020-09-27%20at%205.42.43%20PM.png?dl=0
