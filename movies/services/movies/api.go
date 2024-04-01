package movies

import (
	"context"

	datasource "github.com/jtomasevic/go-graphql-federation-demo/movies/services/movies/data_source"
	"github.com/jtomasevic/go-graphql-federation-demo/movies/services/movies/data_store/model"
	"github.com/jtomasevic/go-graphql-federation-demo/movies/services/movies/data_store/repo"
)

type Service interface {
	GetMovies(ctx context.Context) ([]Movie, error)
	GetMovieByImdbId(ctx context.Context, imdbId string) (Movie, error)
	GetMoviesByImdbId(ctx context.Context, imdbIds []string) ([]Movie, error)
}

type DataStore interface {
	GetMovies(ctx context.Context) ([]model.Movie, error)
	GetMovieByImdbId(ctx context.Context, imdbId string) (model.Movie, error)
	GetMoviesByImdbId(ctx context.Context, imdbIds []string) ([]model.Movie, error)
}

func NewDataStore(dataSource datasource.DataSource) *repo.MovieRepo {
	return &repo.MovieRepo{
		DataSource: dataSource,
	}
}

func NewService(dataStore DataStore) *MovieService {
	return &MovieService{
		dataStore: dataStore,
	}
}
