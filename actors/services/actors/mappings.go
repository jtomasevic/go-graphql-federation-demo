package actors

import (
	"strings"

	"github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors/data_store/model"
)

func dbToServiceActor(actor model.Actor) Actor {
	result := Actor{
		ID:     actor.ID,
		Name:   actor.Name,
		ImdbID: actor.ImdbID,
	}
	movies := strings.Split(actor.MovieImdbs, ";")
	for _, movie := range movies {
		result.MovieImdbIds = append(result.MovieImdbIds, strings.Trim(movie, " "))
	}
	result.MovieImdbIds = movies
	return result
}

func dbToServiceActors(actors []model.Actor) []Actor {
	result := []Actor{}
	for _, actor := range actors {
		result = append(result, dbToServiceActor(actor))
	}
	return result
}
