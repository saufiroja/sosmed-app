package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

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
	err := c.KafkaConsumer.SubscribeTopic(messaging.InsertUserTopic)
	if err != nil {
		c.logger.Error("failed to subscribe to topic")
	}

	run := true

	for run {
		ev := c.KafkaConsumer.KafkaConsumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			// application-specific processing
			c.logger.Info("message received")
			var user models.User
			err := json.Unmarshal(e.Value, &user)
			if err != nil {
				c.logger.Error("failed to unmarshal message")
			}

			ctx := context.Background()
			err = c.userService.InsertUser(ctx, &user)
			if err != nil {
				c.logger.Error("failed to insert user ", err)
			}

			c.logger.Info("user inserted")
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}

	c.KafkaConsumer.Close()
}
