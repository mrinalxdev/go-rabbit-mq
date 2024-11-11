package main

import (
	"log"

	"github.com/mrinalxdev/rabbit-mq-go/rabbitmq"
)


func main() {
	conn := rabbitmq.ConnectRabbitMQ()
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
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
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("Received: %s", msg.Body)
		}
	}()
	log.Println("Waiting for messages. To exit press CTRL+C")
	<-forever
}