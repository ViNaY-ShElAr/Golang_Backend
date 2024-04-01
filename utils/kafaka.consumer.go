package utils

import (
	kf "GO_PROJECT/kafka"
	"GO_PROJECT/model"
	"GO_PROJECT/db/cassandra"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func StartKafkaConsumer() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	go func() {
		<-sig
		cancel() // Cancel context on shutdown signal
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context canceled. Exiting Kafka consumer loop.")
			return // Exit the loop when context is canceled.
		default:
			msg, err := kf.KafkaReader.ReadMessage(ctx)
			if err != nil {
				fmt.Println("Failed to read message: ", err)
				continue // Continue to the next iteration in case of error.
			}
			fmt.Printf("Received message: Key: %s, Value: %s\n", string(msg.Key), string(msg.Value))

			var messageData model.MessageData
			if err := json.Unmarshal(msg.Value, &messageData); err != nil {
				fmt.Println("Failed to unmarshal message: ", err)
				continue // Continue to the next iteration in case of error.
			}

			// Insert data in the database
			cassandra.InsertMessagedata(messageData)
		}
	}
}
