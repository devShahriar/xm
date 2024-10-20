package kafka

import (
	"log"

	"github.com/devShahriar/xm/internal/config"
)

func PublishEvent(message string, topic string) error {
	broker, err := NewKafkaProducer([]string{config.GetAppConfig().KafkaBrokerUrl}, topic)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	return broker.PublishMessage(message)
}
