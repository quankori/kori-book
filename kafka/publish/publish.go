package main

import (
	"log"
	"strconv"

	"github.com/IBM/sarama"
)

func main() {
	// Set up Kafka producer configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Create a Kafka producer
	brokerAddress := "127.0.0.1:9092"
	producer, err := sarama.NewSyncProducer([]string{brokerAddress}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}
	defer producer.Close()
	topic := "kori_test"

	for i := 0; i < 10; i++ {
		// Publish messages to a Kafka topic
		message := "Hello, a!" + strconv.FormatInt(int64(i), 10)

		// Create a Kafka message
		kafkaMessage := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(message),
		}

		// Send the message to Kafka
		_, _, err = producer.SendMessage(kafkaMessage)
		if err != nil {
			log.Fatalf("Error sending message: %v", err)
		}

		log.Printf("Message sent to Kafka: %s", message)
	}
}
