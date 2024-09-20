package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type Message struct {
	ID      string `json:"id"`
	Pattern string `json:"pattern"` // Add the pattern directly in the message payload
	Body    string `json:"body"`
	Data    string `json:"data"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func sendMessageToQueue(message Message) (string, error) {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://user:password@rabbitmq:5672")
	if err != nil {
		return "", fmt.Errorf("Failed to connect to RabbitMQ: %w", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		return "", fmt.Errorf("Failed to open a channel: %w", err)
	}
	defer ch.Close()

	// Declare the main queue
	q, err := ch.QueueDeclare(
		"queue_message", // Queue name
		true,            // Durable
		false,           // Delete when unused
		false,           // Exclusive
		false,           // No-wait
		nil,             // Arguments
	)
	if err != nil {
		return "", fmt.Errorf("Failed to declare a queue: %w", err)
	}

	// Declare a temporary queue for the reply
	replyQueue, err := ch.QueueDeclare(
		"",    // Name (empty means RabbitMQ will generate a unique name)
		false, // Durable
		true,  // Auto-delete (true to auto-delete)
		true,  // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return "", fmt.Errorf("Failed to declare a reply queue: %w", err)
	}

	// Create a correlation ID for this request
	corrId := randomString(32)
	message.ID = corrId

	// Serialize the message to JSON
	body, err := json.Marshal(message)
	if err != nil {
		return "", fmt.Errorf("Failed to encode message to JSON: %w", err)
	}

	// Publish the message with a ReplyTo and CorrelationId
	err = ch.Publish(
		"",     // No exchange
		q.Name, // Queue name (routing key = queue name)
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			ReplyTo:       replyQueue.Name, // This is the queue where we expect the response
			Body:          body,
		})
	if err != nil {
		return "", fmt.Errorf("Failed to publish a message: %w", err)
	}

	// Channel to receive messages
	msgs := make(chan string)

	// Consume the response from the temporary reply queue
	go func() {
		deliveries, err := ch.Consume(
			replyQueue.Name, // Queue name
			"",              // Consumer tag
			true,            // Auto-ack
			false,           // Exclusive
			false,           // No-local
			false,           // No-wait
			nil,             // Arguments
		)
		if err != nil {
			log.Printf("Failed to register a consumer: %s", err)
			return
		}

		// Wait for the response and match the correlation ID
		for d := range deliveries {
			if corrId == d.CorrelationId {
				// We have received the correct response
				log.Printf("Received response: %s", d.Body)
				msgs <- string(d.Body)
				return
			}
		}
	}()

	// Wait for a response with a timeout
	select {
	case response := <-msgs:
		return response, nil
	case <-time.After(10 * time.Second): // Timeout of 10 seconds
		return "", fmt.Errorf("Request timed out waiting for response")
	}
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Define the POST endpoint to send a message to RabbitMQ
	router.GET("/send", func(c *gin.Context) {
		message := Message{
			Pattern: "pattern_one",
			Body:    "Hello from Golang",
			Data:    "test",
			ID:      "",
		}

		// Send the message to RabbitMQ
		_, err := sendMessageToQueue(message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return success response
		c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
	})

	// Run the server on port 8080
	router.Run(":4000")
}
