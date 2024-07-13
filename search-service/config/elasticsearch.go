package config

import "os"

func initElasticsearch(conf *AppConfig) {
	url := os.Getenv("ELASTIC_URL")

	conf.Elasticsearch.URL = url
}
