package logic

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"os4gophers/domain"

	"github.com/opensearch-project/opensearch-go"
)

func LookupMovieTitleByDocumentID(ctx context.Context, opensearchclient *opensearch.Client, arrLength int) {
	documentId, err := rand.Int(rand.Reader, big.NewInt(int64(arrLength)))
	if err != nil {
		panic(err)
	}
	response, err := opensearchclient.Get("movies", documentId.String())
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var getResponse = domain.GetResponse{}
	err = json.NewDecoder(response.Body).Decode(&getResponse)
	if err != nil {
		panic(err)
	}

	movieTitle := getResponse.Source.Title
	fmt.Printf("🟦 Movie with the ID %d: %s \n", documentId.Int64(), movieTitle)
}
