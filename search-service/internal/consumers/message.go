package consumer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/saufiroja/sosmed/search-service/internal/models"
	"github.com/saufiroja/sosmed/search-service/internal/services"
	"github.com/saufiroja/sosmed/search-service/pkg/messaging"
	"github.com/sirupsen/logrus"
)

type Consumer struct {
	KafkaConsumer *messaging.KafkaConsumer
	userService   services.UserServiceInterface
	logger        *logrus.Logger
}

func NewConsumer(kafkaConsumer *messaging.KafkaConsumer, userService services.UserServiceInterface, logger *logrus.Logger) *Consumer {
	return &Consumer{
		KafkaConsumer: kafkaConsumer,
		userService:   userService,
		logger:        logger,
	}
}

func (c *Consumer) Start() {
	c.logger.Info("Starting consumer...")
	go c.startInsertUserConsumer()
}

func (c *Consumer) startInsertUserConsumer() {
	if err := c.KafkaConsumer.SubscribeTopic(messaging.InsertUserTopic); err != nil {
		c.logger.Error("Failed to subscribe to topic: ", err)
		return
	}

	for {
		ev := c.KafkaConsumer.KafkaConsumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			c.processMessage(e)
		case kafka.Error:
			c.logger.Error(fmt.Sprintf("Kafka error: %v", e))
		default:
			c.logger.Debug(fmt.Sprintf("Ignored event: %v", e))
		}
	}
}

func (c *Consumer) processMessage(message *kafka.Message) {
	c.logger.Info("Message received")
	var user models.User
	if err := json.Unmarshal(message.Value, &user); err != nil {
		c.logger.Error("Failed to unmarshal message: ", err)
		return
	}

	ctx := context.Background()
	if err := c.userService.InsertUser(ctx, &user); err != nil {
		c.logger.Error("Failed to insert user: ", err)
		return
	}

	c.logger.Info("User inserted successfully")
}
