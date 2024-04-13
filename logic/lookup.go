package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os4gophers/domain"
	"strconv"
	"time"

	"github.com/opensearch-project/opensearch-go"
)

func QueryMovieByDocumentID(ctx context.Context) {

	movies := ctx.Value(domain.MoviesKey).([]domain.Movie)
	client := ctx.Value(domain.ClientKey).(*opensearch.Client)

	rand.Seed(time.Now().UnixNano())
	documentID := rand.Intn(len(movies) - 1)
	response, err := client.Get("movies", strconv.Itoa(documentID))
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
	fmt.Printf("✅ Movie with the ID %d: %s \n", documentID, movieTitle)

}
