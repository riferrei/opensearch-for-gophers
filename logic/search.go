package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os4gophers/domain"
	"strings"

	"github.com/opensearch-project/opensearch-go"
)

func SearchBestMatrixMovies(ctx context.Context, opensearchClient *opensearch.Client) {
	var searchBuffer bytes.Buffer
	search := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match": map[string]string{
						"actors.en": "keanu reeves",
					},
				},
				"filter": []map[string]interface{}{
					{
						"term": map[string]string{
							"genres.keyword": "Action",
						},
					},
					{
						"range": map[string]interface{}{
							"rating": map[string]float64{
								"gte": 7.0,
							},
						},
					},
					{
						"range": map[string]interface{}{
							"year": map[string]int{
								"gte": 1995,
								"lte": 2005,
							},
						},
					},
				},
			},
		},
	}
	err := json.NewEncoder(&searchBuffer).Encode(search)
	if err != nil {
		panic(err)
	}

	response, err := opensearchClient.Search(
		opensearchClient.Search.WithContext(ctx),
		opensearchClient.Search.WithIndex("movies"),
		opensearchClient.Search.WithBody(&searchBuffer),
		opensearchClient.Search.WithTrackTotalHits(true),
		opensearchClient.Search.WithPretty(),
	)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var searchResponse = domain.SearchResponse{}
	err = json.NewDecoder(response.Body).Decode(&searchResponse)
	if err != nil {
		panic(err)
	}

	if searchResponse.Hits.Total.Value > 0 {
		var movieTitles []string
		for _, movieTitle := range searchResponse.Hits.Hits {
			movieTitles = append(movieTitles, movieTitle.Source.Title)
		}
		fmt.Printf("🟦 Best Matrix movies with Keanu Reeves: [%s] \n", strings.Join(movieTitles, ", "))
	}
}
