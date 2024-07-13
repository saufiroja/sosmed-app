package consumer

import (
	"context"
	"encoding/json"
	"fmt"

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
	err := c.KafkaConsumer.SubscribeTopic(messaging.InsertUserTopic)
	if err != nil {
		c.logger.Error("failed to subscribe to topic")
	}

	for {
		c.logger.Info("waiting for message...")
		msg, err := c.KafkaConsumer.KafkaConsumer.ReadMessage(-1)
		if err != nil {
			c.logger.Error("failed to read message")
		}

		var user models.User
		if err := json.Unmarshal(msg.Value, &user); err != nil {
			c.logger.Error("failed to unmarshal message")
			return
		}

		if err := c.userService.InsertUser(context.Background(), &user); err != nil {
			c.logger.Error(fmt.Sprintf("failed to insert user, err: %v", err))
			return
		}

		c.logger.Info(fmt.Sprintf("user inserted: %v", user))
	}
}
