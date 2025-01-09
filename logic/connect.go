package logic

import (
	"context"
	"os4gophers/domain"

	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

const (
	opensearchEndpoint = "http://localhost:9200"
)

func ConnectWithOpenSearch(ctx context.Context) context.Context {

	newClient, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{
			opensearchEndpoint,
		},
		Username: "opensearch_user",
		Password: "W&lcome123",
	})
	if err != nil {
		panic(err)
	}

	pingRequest := opensearchapi.PingRequest{
		Pretty:     true,
		Human:      true,
		ErrorTrace: true,
	}

	_, err = pingRequest.Do(ctx, newClient)

	if err != nil {
		panic(err)
	}

	return context.WithValue(ctx, domain.ClientKey, newClient)
}
