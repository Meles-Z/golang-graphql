package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// before i process here is what i have to know
// Configure Kafka Writer
// Set the topic you want to write to
// Prepare the Message Payload
// Create Kafka Message
// Send the Message
// Handle Errors

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// 1. Configure Kafka producer
	cfg := kafka.ConfigMap{
		"bootstrap.servers":      "localhost:9092",
		"acks":                   "all",    // Wait for all replicas to acknowledge
		"retries":                3,        // Retry if temporary failure
		"queue.buffering.max.ms": 5,        // Buffer before sending
		"compression.type":       "snappy", // Compress messages
		"go.delivery.reports":    true,     // Enable delivery reports
		"enable.idempotence":     true,     // Ensure no duplicates
	}

	p, err := kafka.NewProducer(&cfg)
	if err != nil {
		fmt.Printf("❌ Failed to create Kafka producer: %s\n", err)
		return
	}
	defer p.Close()

	// 2. Setup delivery report handler (best practice)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("❌ Delivery failed: %v\n", ev.TopicPartition.Error)
				} else {
					fmt.Printf("✅ Message delivered to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// 3. Prepare message payload
	per := Person{
		Name:  "Meles",
		Email: "meles@example.com",
	}

	msgBytes, err := json.Marshal(per)
	if err != nil {
		fmt.Printf("❌ Failed to marshal message: %s\n", err)
		return
	}

	// 4. Create Kafka message and send
	topic := "my-topic"
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msgBytes,
		Key:            []byte(per.Email), // optional: good for partitioning
	}, nil)

	if err != nil {
		fmt.Printf("❌ Failed to produce message: %s\n", err)
		return
	}

	// 5. Wait for message delivery or handle graceful shutdown
	// (Don't exit until all messages are delivered)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("📨 Message sent, waiting for delivery...")
	<-sigChan

	fmt.Println("📦 Flushing remaining messages...")
	p.Flush(15_000) // wait up to 15 seconds for delivery before exit
}
