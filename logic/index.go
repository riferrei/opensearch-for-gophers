package logic

import (
	"context"
	"es4gophers/domain"
	"fmt"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

func IndexMoviesAsDocuments(ctx context.Context) {

	movies := ctx.Value(domain.MoviesKey).([]domain.Movie)
	client := ctx.Value(domain.ClientKey).(*elasticsearch.Client)

	/*
		for documentID, document := range movies {
			res, err := client.Index("movies", esutil.NewJSONReader(document),
				client.Index.WithDocumentID(strconv.Itoa(documentID)))
			if err == nil {
				fmt.Println(res)
			} else {
				fmt.Println(err)
			}
		}
	*/

	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:      "movies",
		Client:     client,
		NumWorkers: 5,
	})
	if err != nil {
		panic(err)
	}

	for documentID, document := range movies {
		err = bulkIndexer.Add(
			ctx,
			esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: strconv.Itoa(documentID),
				Body:       esutil.NewJSONReader(document),
			},
		)
		if err != nil {
			panic(err)
		}
	}

	bulkIndexer.Close(ctx)
	biStats := bulkIndexer.Stats()
	fmt.Printf("✅ Movies indexed on Elasticsearch: %d \n", biStats.NumIndexed)

}
