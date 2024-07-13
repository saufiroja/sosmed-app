package config

import "os"

func initKafkaClient(conf *AppConfig) {
	url := os.Getenv("KAFKA_BROKER")

	conf.KafkaClient.URL = url
}
