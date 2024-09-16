package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://user:password@rabbitmq:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare a custom exchange
	exchangeName := "my_exchange"
	err = ch.ExchangeDeclare(
		exchangeName, // name of the exchange
		"direct",     // type (direct exchange routes based on routing key)
		true,         // durable
		false,        // auto-delete when unused
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	// Declare the queue
	queueName := "nestjs_queue"
	_, err = ch.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Bind the queue to the custom exchange with routing key "test"
	err = ch.QueueBind(
		queueName,    // queue name
		"test",       // routing key
		exchangeName, // custom exchange name
		false,
		nil,
	)
	failOnError(err, "Failed to bind the queue")

	// Publish a message with routing key "test" to the custom exchange
	body := "Hello from Golang!"
	err = ch.Publish(
		exchangeName, // publish to custom exchange
		"test",       // routing key (matches @MessagePattern("test") in NestJS)
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	fmt.Println(" [x] Sent message:", body)

	// Keep the program alive until an interrupt signal is received
	forever := make(chan bool)

	// Handle interrupt signals to gracefully exit the program
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Printf("Received signal: %s. Exiting...\n", sig)
		close(forever)
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever // Blocks the program from exiting
}
