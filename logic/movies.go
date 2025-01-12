package logic

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os4gophers/domain"
	"sync"
)

func LoadMoviesFromFile(fileName string) ([]domain.Movie, error) {
	const (
		concurrency = 5
	)
	var (
		movies    []domain.Movie
		waitGroup = new(sync.WaitGroup)
		workQueue = make(chan string)
		mutex     = &sync.Mutex{}
		errChan   = make(chan error, concurrency)
	)

	go func() {
		moviesFile, err := os.Open(fileName)
		if err != nil {
			errChan <- err
			close(workQueue)
			return
		}
		defer func(moviesFile *os.File) {
			err := moviesFile.Close()
			if err != nil {
				errChan <- err
			}
		}(moviesFile)
		scanner := bufio.NewScanner(moviesFile)
		for scanner.Scan() {
			workQueue <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			errChan <- err
		}
		close(workQueue)
	}()

	for i := 0; i < concurrency; i++ {
		waitGroup.Add(1)
		go func(workQueue chan string, waitGroup *sync.WaitGroup) {
			defer waitGroup.Done()
			for entry := range workQueue {
				movieRaw := domain.MovieRaw{}
				err := json.Unmarshal([]byte(entry), &movieRaw)
				if err != nil {
					errChan <- err
					continue
				}
				movie := func(movieRaw domain.MovieRaw) domain.Movie {
					return domain.Movie{
						Title:       movieRaw.Title,
						Year:        movieRaw.Year,
						RunningTime: movieRaw.Info.RunningTime,
						ReleaseDate: movieRaw.Info.ReleaseDate,
						Rating:      movieRaw.Info.Rating,
						Genres:      movieRaw.Info.Genres,
						Actors:      movieRaw.Info.Actors,
						Directors:   movieRaw.Info.Directors,
					}
				}(movieRaw)
				mutex.Lock()
				movies = append(movies, movie)
				mutex.Unlock()
			}
		}(workQueue, waitGroup)
	}

	waitGroup.Wait()
	close(errChan)

	// Collect errors
	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		errorMsg := "Errors occurred while loading movies:"
		for _, err := range errs {
			errorMsg += "\n- " + err.Error()
		}
		return movies, errors.New(errorMsg)
	}

	fmt.Printf("🟦 Movies loaded from file: %d \n", len(movies))
	return movies, nil
}
