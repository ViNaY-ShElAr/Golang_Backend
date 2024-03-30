package utils

import (
	"fmt"
	"context"
	"encoding/json"

	"GO_PROJECT/model"

	"github.com/segmentio/kafka-go"
)

func StartKafkaProducer(KafkaWriter *kafka.Writer, message model.MessageData) error {

	data, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	err = KafkaWriter.WriteMessages(context.Background(), kafka.Message{
		Value: data,
	})
	return err

}
