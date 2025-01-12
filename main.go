package main

import (
	"context"
	"log"
	"os4gophers/logic"
)

func main() {
	ctx := context.Background()

	opensearchClient, err := logic.ConnectWithOpenSearch(ctx)
	if err != nil {
		log.Fatalf("Error connecting with OpenSearch: %v", err)
	}

	movies, err := logic.LoadMoviesFromFile("movies.json")
	if err != nil {
		log.Fatalf("Error loading movies from file: %v", err)
	}

	logic.IndexMoviesAsDocuments(ctx, opensearchClient, movies)
	logic.LookupMovieTitleByDocumentID(ctx, opensearchClient, len(movies))
	logic.MovieCountPerGenreAgg(ctx, opensearchClient)
	logic.SearchBestMatrixMovies(ctx, opensearchClient)
}
