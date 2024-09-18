package main

import (
	"fmt"
	"log"

	"encoding/json"

	"github.com/streadway/amqp"
)

type Message struct {
	Pattern string `json:"pattern"` // Add the pattern directly in the message payload
	Body    string `json:"body"`
	Data    string `json:"data"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func sendMessageToQueue(message Message) {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://user:password@rabbitmq:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare the queue
	q, err := ch.QueueDeclare(
		"queue_message", // Queue name (same as in NestJS)
		true,            // Durable
		false,           // Delete when unused
		false,           // Exclusive
		false,           // No-wait
		nil,             // Arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Serialize the message to JSON
	body, err := json.Marshal(message)
	failOnError(err, "Failed to encode message to JSON")

	// Publish the message to the queue
	err = ch.Publish(
		"",     // No exchange
		q.Name, // Queue name (routing key = queue name)
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")

	fmt.Printf("Message sent with pattern %s: %s\n", message.Pattern, string(body))
}

func main() {
	// Pattern 1 Message
	msg1 := Message{
		Pattern: "pattern_one", // Set the pattern directly in the payload
		Body:    "Hello from Golang to pattern_1",
		Data:    "test",
	}
	sendMessageToQueue(msg1)

	// Pattern 2 Message
	msg2 := Message{
		Pattern: "pattern_two", // Set the pattern directly in the payload
		Body:    "Hello from Golang to pattern_2",
		Data:    "22",
	}
	sendMessageToQueue(msg2)
}
