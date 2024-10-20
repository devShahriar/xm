package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

var KafkaBroker *KafkaProducer

func GetKafkaBroker() *KafkaProducer {
	return KafkaBroker
}

type KafkaProducer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewKafkaProducer(brokers []string, topic string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Printf("kafka producer error %v", err)
		return nil, err
	}

	return &KafkaProducer{
		producer: producer,
		topic:    topic,
	}, nil
}

func (p *KafkaProducer) PublishMessage(message string) error {
	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	} else {
		log.Printf("Message sent successfully to topic %v partition %d at offset %d\n", p.topic, partition, offset)
	}

	log.Printf("Message is stored in partition %d, offset %d\n", partition, offset)
	p.Close()
	return nil
}

func (p *KafkaProducer) Close() error {
	return p.producer.Close()
}
