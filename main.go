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
	logic.LookupMovieTitleByMovieID(ctx)
	logic.SearchBestMatrixMovies(ctx)
	logic.MovieCountPerGenreAgg(ctx)
}
