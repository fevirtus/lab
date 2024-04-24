package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://admin:123@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare a queue
	queue, err := ch.QueueDeclare(
		"node-test", // Queue name
		true,        // Durable
		false,       // Delete when unused
		false,       // Exclusive
		false,       // No-wait
		nil,         // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Bind the queue to the exchange
	err = ch.QueueBind(
		queue.Name,        // tên của hàng đợi
		"name of cluster", // routing key
		"event",           // tên của Exchange
		false,
		nil,
	)

	// Consume messages from the queue
	msgs, err := ch.Consume(
		queue.Name, // Queue name
		"",         // Consumer name
		true,       // Auto-acknowledge
		false,      // Exclusive
		false,      // No-local
		false,      // No-wait
		nil,        // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Start consuming messages
	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			fmt.Printf("Received message: %s\n", msg.Body)
		}
	}()

	log.Println("Consumer started. Press Ctrl+C to exit.")
	<-forever
}
