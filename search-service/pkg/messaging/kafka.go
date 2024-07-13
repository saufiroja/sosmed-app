package messaging

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/saufiroja/sosmed/search-service/config"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	KafkaConsumer *kafka.Consumer
	Logger        *logrus.Logger
}

func NewKafkaConsumer(conf *config.AppConfig, logger *logrus.Logger) *KafkaConsumer {
	KafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": conf.KafkaClient.URL,
		"group.id":          "golang",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(KafkaConfig)
	if err != nil {
		logger.Panicf("Failed to create consumer: %s\n", err)
	}

	return &KafkaConsumer{
		KafkaConsumer: consumer,
		Logger:        logger,
	}
}

func (k *KafkaConsumer) SubscribeTopic(topic string) error {
	err := k.KafkaConsumer.Subscribe(topic, nil)
	if err != nil {
		k.Logger.Errorf("Failed to subscribe to topic %s: %s\n", topic, err)
		return err
	}

	return nil
}

func (k *KafkaConsumer) Close() error {
	return k.KafkaConsumer.Close()
}
