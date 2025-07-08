package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"}, // <-- use localhost here
		Topic:   "pg-users",
		GroupID: "go-consumer-group",
	})

	fmt.Println("📥 Listening for messages on topic 'pg-users'...")

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("📨 Received: %s\n", msg.Value)
	}
}
