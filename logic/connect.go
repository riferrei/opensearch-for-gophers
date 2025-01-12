package logic

import (
	"context"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"log"
	"os"
)

var opensearchConnectionURL = func() string {
	if osConnURL := os.Getenv("OPENSEARCH_CONNECTION_URL"); osConnURL != "" {
		return osConnURL
	}
	return "http://localhost:9200"
}()

func ConnectWithOpenSearch(ctx context.Context) (*opensearch.Client, error) {
	newClient, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{
			opensearchConnectionURL,
		},
		Username: "opensearch_user",
		Password: "W&lcome123",
	})
	if err != nil {
		log.Printf("Error parsing the Redis URL: %v", err)
		return nil, err
	}

	pingRequest := opensearchapi.PingRequest{
		Pretty:     true,
		Human:      true,
		ErrorTrace: true,
	}

	_, err = pingRequest.Do(ctx, newClient)
	if err != nil {
		log.Printf("Ping request failed: %v", err)
		return nil, err
	}

	return newClient, nil
}
