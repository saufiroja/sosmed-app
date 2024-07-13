package database

import (
	esv8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/saufiroja/sosmed/search-service/config"
	"github.com/sirupsen/logrus"
)

type Elasticsearch struct {
	Client *esv8.Client
}

func NewElasticsearch(conf *config.AppConfig, logger *logrus.Logger) *Elasticsearch {
	// addresses := fmt.Sprintf("%s:%s", conf.Elasticsearch.Host, conf.Elasticsearch.Port)
	client, err := esv8.NewClient(
		esv8.Config{
			Addresses: []string{
				conf.Elasticsearch.URL,
			},
		},
	)

	if err != nil {
		logger.Panicf("Error creating the client: %s", err)
	}

	logger.Info("connected to Elasticsearch")

	return &Elasticsearch{
		Client: client,
	}
}

func (e *Elasticsearch) DBConnecion() *esv8.Client {
	return e.Client
}
