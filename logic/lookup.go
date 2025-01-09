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

func LookupMovieTitleByMovieID(ctx context.Context) {

	movies := ctx.Value(domain.MoviesKey).([]domain.Movie)
	client := ctx.Value(domain.ClientKey).(*opensearch.Client)

	documentId, err := rand.Int(rand.Reader, big.NewInt(int64(len(movies))))
	if err != nil {
		panic(err)
	}
	response, err := client.Get("movies", documentId.String())
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
	fmt.Printf("🟦 Movie with the ID '%d': %s \n", documentId.Int64(), movieTitle)
}
