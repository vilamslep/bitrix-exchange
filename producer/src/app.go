package main

import (
	"context"
	"log"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	fmt.Println("Hello from Docker")
	// Producer()
}

func Producer() {

	// to produce messages
	topic := "order"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "172.20.0.3:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{
			Key:   []byte("order"),
			Value: []byte("one!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}
