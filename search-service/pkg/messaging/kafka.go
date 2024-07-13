package messaging

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/saufiroja/sosmed/search-service/config"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	KafkaConsumer *kafka.Consumer
}

func NewKafkaConsumer(conf *config.AppConfig, logger *logrus.Logger) *KafkaConsumer {
	KafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": conf.KafkaClient.URL,
		"group.id":          "golang",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(KafkaConfig)
	if err != nil {
		logger.Error("failed to connect to Kafka")
	}

	logger.Info("connected to Kafka")

	return &KafkaConsumer{
		KafkaConsumer: consumer,
	}
}

func (k *KafkaConsumer) SubscribeTopic(topic string) error {
	err := k.KafkaConsumer.Subscribe(topic, nil)
	if err != nil {
		return err
	}

	return nil
}

func (k *KafkaConsumer) Close(ctx context.Context) error {
	return k.KafkaConsumer.Close()
}
