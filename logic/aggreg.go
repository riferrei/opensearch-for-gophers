package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os4gophers/domain"

	"github.com/opensearch-project/opensearch-go"
)

func MovieCountPerGenreAgg(ctx context.Context) {

	client := ctx.Value(domain.ClientKey).(*opensearch.Client)

	var searchBuffer bytes.Buffer
	aggregRequest := domain.AggregationRequest{
		Size: 0,
		Aggs: &domain.Aggs{
			MovieCountPerGenre: &domain.MovieCountPerGenreRequest{
				Terms: &domain.Terms{
					Field: "genres.keyword",
					Size:  5,
				},
			},
		},
	}
	err := json.NewEncoder(&searchBuffer).Encode(aggregRequest)
	if err != nil {
		panic(err)
	}

	response, err := client.Search(
		client.Search.WithContext(ctx),
		client.Search.WithIndex("movies"),
		client.Search.WithBody(&searchBuffer),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var aggregResponse = domain.AggregationResponse{}
	err = json.NewDecoder(response.Body).Decode(&aggregResponse)
	if err != nil {
		panic(err)
	}

	if len(aggregResponse.Aggregations.MovieCountPerGenreResponse.Buckets) > 0 {
		fmt.Printf("✅ Top 5 Genres and their Movie Count: \n")
		for _, bucket := range aggregResponse.Aggregations.MovieCountPerGenreResponse.Buckets {
			fmt.Printf("   🚀 %s = %d\n", bucket.Key, bucket.DocumentCount)
		}
	}

}
