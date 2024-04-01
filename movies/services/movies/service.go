package movies

import (
	"context"
)

type MovieService struct {
	dataStore DataStore
}

func(service *MovieService) GetMovies(ctx context.Context) ([]Movie, error) {
	movies, err:= service.dataStore.GetMovies(ctx)
	if err != nil {
		return nil, err
	}
	result:= []Movie{}
	for _, movie := range movies {
		result = append(result, Movie(movie)) // jt note: implicit conversion.
	}
	return result, nil
}

func (service *MovieService) GetMovieByImdbId(ctx context.Context, imdbId string) (Movie, error) {
	movie, err:= service.dataStore.GetMovieByImdbId(ctx, imdbId)
	if err != nil {
		return Movie{}, err
	}
	return Movie(movie), nil
}

func (service *MovieService) GetMoviesByImdbId(ctx context.Context, imdbIds []string) ([]Movie, error) {
	movies, err:= service.dataStore.GetMoviesByImdbId(ctx, imdbIds)
	if err != nil {
		return nil, err
	}
	result:= []Movie{}
	for _, movie := range movies {
		result = append(result, Movie(movie)) // jt note: implicit conversion.
	}
	return result, nil
}