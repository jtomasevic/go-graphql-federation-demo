package resolvers

import (
	"github.com/jtomasevic/go-graphql-federation-demo/movies/graph/model"
	"github.com/jtomasevic/go-graphql-federation-demo/movies/services/movies"
)

func fromServiceToGqlMovies(movies []movies.Movie) []model.Movie {
	result := []model.Movie{}
	for _, movie := range movies {
		result = append(result, model.Movie(movie))
	}
	return result
}
