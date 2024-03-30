package utils

import (
	"GO_PROJECT/model"
	kf "GO_PROJECT/kafka"
	"context"
	"encoding/json"
	"fmt"
)

func StartKafkaConsumer() {
	for {
		msg, err := kf.KafkaReader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Failed to read message: ", err)
		}
		fmt.Printf("Received message: Key: %s, Value: %s\n", string(msg.Key), string(msg.Value))

		var messageData model.MessageData
		if err := json.Unmarshal(msg.Value, &messageData); err != nil {
			fmt.Println("Failed to unmarshal message: ", err)
			continue
		}
		
		// insert data in db

	}
}
