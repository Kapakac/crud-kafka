package producer

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"time"
)

const partition = 0

func WriteMessages(messages map[string]interface{}) {
	host := os.Getenv("KAFKA_HOST")
	topic := os.Getenv("KAFKA_TOPIC")

	conn, err := kafka.DialLeader(context.Background(), "tcp", host, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	err = conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		log.Fatal("failed to set deadline:", err)
	}

	for key, value := range messages {
		entity, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		msg := kafka.Message{
			Key:   []byte(key),
			Value: entity,
		}
		_, err = conn.WriteMessages(msg)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
