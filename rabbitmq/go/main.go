package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Function to publish message to RabbitMQ and wait for the response
func publishAndWaitForResponse(message string) (string, error) {
	conn, err := amqp.Dial("amqp://user:password@rabbitmq:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare a reply queue to get the response back
	replyQueue, err := ch.QueueDeclare(
		"",    // Name of the reply queue (empty string lets RabbitMQ generate a unique name)
		false, // Durable
		false, // Delete when unused
		true,  // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	failOnError(err, "Failed to declare a reply queue")

	// Set up a channel to consume the response
	msgs, err := ch.Consume(
		replyQueue.Name, // Queue name
		"",              // Consumer tag
		true,            // Auto-ack
		false,           // Exclusive
		false,           // No-local
		false,           // No-wait
		nil,             // Args
	)
	failOnError(err, "Failed to register a consumer")

	// Generate a unique correlation ID
	corrID := randomString(32)

	// Publish the message with the correlation ID and reply queue
	err = ch.Publish(
		"",             // Exchange
		"send_message", // Routing key
		false,          // Mandatory
		false,          // Immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrID,
			ReplyTo:       replyQueue.Name, // Reply-to header with reply queue
			Body:          []byte(message),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", message)

	// Wait for a response on the reply queue
	for msg := range msgs {
		if msg.CorrelationId == corrID {
			// Return the response message if the correlation ID matches
			log.Printf(" [x] Received response: %s", msg.Body)
			return string(msg.Body), nil
		}
	}

	return "", fmt.Errorf("No response received")
}

// Utility function to generate a random string for correlation IDs
func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	r := gin.Default()

	r.GET("/send", func(c *gin.Context) {
		message := "Hello from Gin to NestJS!"
		response, err := publishAndWaitForResponse(message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": response})
	})

	r.Run(":4000")
}
