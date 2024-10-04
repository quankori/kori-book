package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

func main() {
	// Set up Kafka consumer configuration
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create a Kafka consumer
	brokerAddress := "127.0.0.1:9092"
	consumer, err := sarama.NewConsumer([]string{brokerAddress}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}
	defer consumer.Close()

	// Subscribe to Kafka topic(s)
	topic := "kori_test"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}
	defer partitionConsumer.Close()

	// Set up a signal channel to gracefully stop the consumer
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Consume and process messages
	for {
		select {
		case message := <-partitionConsumer.Messages():
			log.Printf("Received message: %s", string(message.Value))
		case err := <-partitionConsumer.Errors():
			log.Printf("Error while consuming: %v", err)
		case <-signals:
			log.Println("Interrupt signal received, shutting down...")
			return
		}
	}
}
