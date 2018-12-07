package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	broker := "kafka1"
	topic := "test"
	groupid := "consumer-group"

	config := &kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupid,
		"auto.offset.reset": "latest",
	}

	// Create the consumer
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatal("Could not create consumer")
	}

	// Subscribe to topic and consume messages
	// topic := getEnv("TOPIC", "test")
	consumer.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	consumer.Close()
}
