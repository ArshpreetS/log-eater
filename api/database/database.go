package database

import (
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func CreateClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		APIKey: os.Getenv("API_KEY"),
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return es, nil
}
