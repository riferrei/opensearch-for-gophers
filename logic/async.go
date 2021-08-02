package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"es4gophers/domain"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

/* This example shows how to use the Async API. Still,
   it wasn't incorporated in the example because it executes
   too darn fast to return a result that is still being
   processed. But the idea is here. */
func BestKeanuActionMoviesAsync(ctx context.Context) {

	client := ctx.Value(domain.ClientKey).(*elasticsearch.Client)

	var searchBuffer bytes.Buffer
	search := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []map[string]interface{}{
					{
						"range": map[string]interface{}{
							"Year": map[string]int{
								"gte": 1900,
								"lte": 2021,
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

	var numberOfDocs int = 0
	asyncRequest := esapi.AsyncSearchSubmitRequest{
		Index:          []string{"movies"},
		Body:           &searchBuffer,
		TrackTotalHits: true,
		Pretty:         true,
		Size:           &numberOfDocs,
	}
	response, err := asyncRequest.Do(ctx, client.Transport)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

}
