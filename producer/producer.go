package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mrinalxdev/rabbit-mq-go/rabbitmq"
	"github.com/streadway/amqp"
)

func main () {
	conn := rabbitmq.ConnectRabbitMQ()
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel : %v", err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"test_queue", 
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue : %v", err)
	}

	for i := 0; i < 10; i ++ {
		message := fmt.Sprintf("Message #%d", i + 1)
		err = ch.Publish(
			"",
			queue.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body: []byte(message),
			},
		)
		if err != nil {
			log.Fatalf("Failed to publish a message : %v", err)
		}
		log.Printf("Sent : %s", message)
		time.Sleep(1 * time.Second)
	}
}