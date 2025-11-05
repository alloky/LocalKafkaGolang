package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// Kafka broker addresses
	brokers := []string{"localhost:9092", "localhost:9093", "localhost:9094"}
	topic := "test-topic"

	// Create producer configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all replicas to acknowledge
	config.Producer.Retry.Max = 3

	// Create producer instance
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	// Produce messages
	log.Printf("Starting to produce messages to topic: %s\n", topic)

	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Hello Kafka! Message #%d at %s", i+1, time.Now().Format(time.RFC3339))
		key := fmt.Sprintf("key-%d", i+1)

		msg := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(key),
			Value: sarama.StringEncoder(message),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("Error producing message: %v\n", err)
		} else {
			log.Printf("Successfully produced message to topic %s [%d] at offset %d: %s\n",
				topic, partition, offset, message)
		}

		// Small delay between messages
		time.Sleep(500 * time.Millisecond)
	}

	log.Println("Producer finished!")
}
