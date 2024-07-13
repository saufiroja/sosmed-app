package messaging

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/saufiroja/sosmed-app/auth-service/config"
	"github.com/sirupsen/logrus"
)

type KafkaProducer struct {
	producer *kafka.Producer
	logger   *logrus.Logger
}

// NewKafkaProducer creates a new KafkaProducer instance
func NewKafkaProducer(conf *config.AppConfig, logger *logrus.Logger) *KafkaProducer {
	fmt.Println("KafkaProducer", conf.KafkaBroker.URL)
	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": conf.KafkaBroker.URL,
	}

	producer, err := kafka.NewProducer(kafkaConfig)
	if err != nil {
		logger.Panicf("failed to create Kafka producer: %s", err)
	}

	// check if the producer was created successfully
	logger.Info("connected to Kafka")

	return &KafkaProducer{
		producer: producer,
		logger:   logger,
	}
}

// Publish sends a message to a specified Kafka topic
func (k *KafkaProducer) Publish(topic string, message []byte) {
	k.logger.Infof("publishing message to topic %s", topic)

	deliveryChan := make(chan kafka.Event, 1)
	defer close(deliveryChan)

	err := k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, deliveryChan)
	if err != nil {
		k.logger.Errorf("failed to publish message: %s", err)
	}

	event := <-deliveryChan
	msg := event.(*kafka.Message)

	if msg.TopicPartition.Error != nil {
		k.logger.Errorf("delivery failed: %v", msg.TopicPartition.Error)
	}

	k.logger.Infof("message published to topic %s", topic)

	k.producer.Flush(15 * 1000)
	k.Close()
}

// Close cleans up the Kafka producer
func (k *KafkaProducer) Close() {
	k.producer.Close()
}
