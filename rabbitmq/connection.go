package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)


func ConnectRabbitMQ() *amqp.Connection  {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ : %v", err)
	}
	return conn
}