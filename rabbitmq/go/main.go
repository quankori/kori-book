package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type Message struct {
	Pattern string `json:"pattern"` // Pattern in the message payload
	Body    string `json:"body"`
	Data    string `json:"data"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func randomCorrelationID() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Int())
}

// sendMessageWithResponse sends a message to RabbitMQ and waits for a response
func sendMessageWithResponse(message Message, responseChan chan<- string) {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://user:password@rabbitmq:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare the main queue (where the message will be sent)
	q, err := ch.QueueDeclare(
		"queue_message", // Queue name (must match the one in NestJS)
		true,            // Durable
		false,           // Delete when unused
		false,           // Exclusive
		false,           // No-wait
		nil,             // Arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Declare a reply queue (temporary queue for receiving responses)
	replyQueue, err := ch.QueueDeclare(
		"",    // Empty name, RabbitMQ will create a random name for this temporary queue
		false, // Non-durable
		false, // Auto-delete when unused
		true,  // Exclusive to this connection
		false, // No-wait
		nil,   // Arguments
	)
	failOnError(err, "Failed to declare a reply queue")

	// Create a channel to consume responses from the reply queue
	messages, err := ch.Consume(
		replyQueue.Name, // The reply queue
		"",              // Consumer tag
		true,            // Auto-ack (can be set to false to manually ack)
		false,           // Exclusive
		false,           // No-local
		false,           // No-wait
		nil,             // Args
	)
	failOnError(err, "Failed to register a consumer")

	// Generate a unique correlation ID
	correlationID := randomCorrelationID()

	// Serialize the message to JSON
	body, err := json.Marshal(message)
	failOnError(err, "Failed to encode message to JSON")

	// Publish the message with reply-to and correlation ID
	err = ch.Publish(
		"",     // No exchange
		q.Name, // Routing key (queue name)
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			ContentType:   "application/json",
			Body:          body,
			ReplyTo:       replyQueue.Name, // Set the callback queue for the response
			CorrelationId: correlationID,   // Unique correlation ID
		})
	failOnError(err, "Failed to publish a message")

	fmt.Printf("Message sent with pattern %s: %s\n", message.Pattern, string(body))

	// Wait for the response in a goroutine
	go func() {
		for d := range messages {
			if d.CorrelationId == correlationID { // Check the correlation ID to match the response
				// Send the response to the response channel
				responseChan <- fmt.Sprintf("Received response for correlation ID %s: %s", correlationID, d.Body)
				break
			}
		}
	}()
}

func main() {
	// Set up the Gin router
	r := gin.Default()

	// POST endpoint to trigger sending a message to RabbitMQ
	r.GET("/send", func(c *gin.Context) {
		msg1 := Message{
			Pattern: "pattern_one", // Set the pattern directly in the payload
			Body:    "Hello from Golang to pattern_1",
			Data:    "test",
		}

		// Create a channel to receive the response asynchronously
		responseChan := make(chan string)

		// Send the message in a goroutine
		sendMessageWithResponse(msg1, responseChan)

		// Immediately respond to the HTTP request while processing the RabbitMQ message asynchronously
		c.JSON(http.StatusAccepted, gin.H{"message": "Message sent, waiting for response..."})

		// Retrieve the response from the goroutine and log it when received
		go func() {
			response := <-responseChan
			fmt.Println(response) // Print the response when it is received
		}()
	})

	// Run the server
	r.Run(":4000")
}
