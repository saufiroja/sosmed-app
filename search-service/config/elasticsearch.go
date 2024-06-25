package config

import "os"

func initElasticsearch(conf *AppConfig) {
	host := os.Getenv("ELASTIC_HOST")
	port := os.Getenv("ELASTIC_PORT")

	conf.Elasticsearch.Host = host
	conf.Elasticsearch.Port = port
}
