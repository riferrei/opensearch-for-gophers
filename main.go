package main

import (
	"context"
	"os4gophers/logic"
)

func main() {

	ctx := context.Background()

	ctx = logic.LoadMoviesFromFile(ctx)
	ctx = logic.ConnectWithOpenSearch(ctx)
	logic.IndexMoviesAsDocuments(ctx)
	logic.QueryMovieByDocumentID(ctx)
	logic.BestKeanuActionMovies(ctx)
	logic.MovieCountPerGenreAgg(ctx)

}
