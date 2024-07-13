package config

import "os"

func initKafkaClient(conf *AppConfig) {
	url := os.Getenv("KAFKA_BROKER")

	conf.KafkaBroker.URL = url
}
