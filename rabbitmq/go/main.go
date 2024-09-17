package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func publishMessage(ch *amqp.Channel, routingKey string, body string) {
	err := ch.Publish(
		"",              // exchange
		"queue_message", // routing key (queue name)
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
			Headers: amqp.Table{
				"pattern": routingKey, // Set the pattern header
			},
		})
	failOnError(err, "Failed to publish a message")
	fmt.Printf(" [x] Sent %s to pattern %s\n", body, routingKey)
}

func main() {
	conn, err := amqp.Dial("amqp://user:password@rabbitmq:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Publish to pattern_one
	publishMessage(ch, "pattern_one", `{"message": "Hello Pattern One"}`)

	// Publish to pattern_two
	publishMessage(ch, "pattern_two", `{"message": "Hello Pattern Two"}`)
}
