package repo

import (
	"context"
	"testing"

	datasource "github.com/jtomasevic/go-graphql-federation-demo/movies/services/movies/data_source"
	"github.com/stretchr/testify/require"
)

func TestMovieRepo_GetMovies(t *testing.T) {
	prepareDb()
	repo := MovieRepo{
		DataSource: datasource.NewDataSource(),
	}
	movies, err := repo.GetMovies(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, movies)
	// Add assertions for the 'movies' variable if needed
}

func prepareDb() {
	InitDb()
	PopulateDb()
}

func TestMovieRepo_GetMovieByImdbId(t *testing.T) {
	prepareDb()
	repo := MovieRepo{
		DataSource: datasource.NewDataSource(),
	}

	movie, err := repo.GetMovieByImdbId(context.Background(), "tt1152836")
	require.NoError(t, err)
	require.NotNil(t, movie)
}

func TestMovieRepo_GetMoviesByImdbId(t *testing.T) {

}
