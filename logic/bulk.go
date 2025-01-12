package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os4gophers/domain"
	"strconv"

	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchutil"
)

func IndexMoviesAsDocuments(ctx context.Context, opensearchClient *opensearch.Client, movies []domain.Movie) {
	bulkIndexer, err := opensearchutil.NewBulkIndexer(opensearchutil.BulkIndexerConfig{
		Index:      "movies",
		Client:     opensearchClient,
		NumWorkers: 5,
	})
	if err != nil {
		panic(err)
	}

	for documentID, document := range movies {
		var buffer bytes.Buffer
		json.NewEncoder(&buffer).Encode(document)
		body := bytes.NewReader(buffer.Bytes())
		err = bulkIndexer.Add(
			ctx,
			opensearchutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: strconv.Itoa(documentID),
				Body:       body,
			},
		)
		if err != nil {
			panic(err)
		}
	}

	bulkIndexer.Close(ctx)
	biStats := bulkIndexer.Stats()
	fmt.Printf("🟦 Movies stored on OpenSearch: %d \n", biStats.NumIndexed)
}
