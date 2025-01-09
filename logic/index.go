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

func IndexMoviesAsDocuments(ctx context.Context) {

	movies := ctx.Value(domain.MoviesKey).([]domain.Movie)
	client := ctx.Value(domain.ClientKey).(*opensearch.Client)

	// for documentID, document := range movies {
	// 	res, err := client.Index("movies", opensearchutil.NewJSONReader(document),
	// 		client.Index.WithDocumentID(strconv.Itoa(documentID)))
	// 	if err == nil {
	// 		fmt.Println(res)
	// 	} else {
	// 		fmt.Println(err)
	// 	}
	// }

	bulkIndexer, err := opensearchutil.NewBulkIndexer(opensearchutil.BulkIndexerConfig{
		Index:      "movies",
		Client:     client,
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
	fmt.Printf("🟦 Movies indexed on OpenSearch: %d \n", biStats.NumIndexed)

}
