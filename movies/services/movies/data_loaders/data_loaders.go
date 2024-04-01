package dataloaders

import (
	"context"

	"net/http"
	"time"

	"github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors"
	"github.com/jtomasevic/go-graphql-federation-demo/movies/services/movies"
)

type ctxKeyType struct{ name string }

var ctxKey = ctxKeyType{"userCtx"}

type Loaders struct {
	ActorsByImdbs *ActorsLoader
	MoviesByImdbs *MoviesLoader
}

func NewLoaders(ctx context.Context, actorService actors.Service, movieService movies.Service) (Loaders, context.Context) {
	dataLoaders := Loaders{}

	// how long to done before sending a batch
	wait := 300 * time.Microsecond

	dataLoaders.ActorsByImdbs = &ActorsLoader{
		wait:     wait,
		maxBatch: 100,
		fetch: func(keys []string) ([]actors.Actor, []error) {
			actorArr, err := actorService.GetActorsByImdbIDs(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}
			result := make([]actors.Actor, len(actorArr))
			actorMap := toActorsMap(actorArr)
			for i, key := range keys {
				result[i] = actorMap[key]
			}
			return result, nil
		},
	}

	dataLoaders.MoviesByImdbs = &MoviesLoader{
		wait:     wait,
		maxBatch: 100,
		fetch: func(keys []string) ([][]movies.Movie, []error) {
			movieArr, err := movieService.GetMoviesByImdbId(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}
			result := make([][]movies.Movie, len(keys))
			for i, key := range keys {
				for _, movie := range movieArr {
					if movie.ImdbID == key {
						result[i] = append(result[i], movie)
					}
				}
			}
			return result, nil
		},
	}

	dataLoadersContext := context.WithValue(ctx, ctxKey, dataLoaders)
	return dataLoaders, dataLoadersContext
}

func GetLoaders(ctx context.Context) Loaders {
	return ctx.Value(ctxKey).(Loaders)
}

func toActorsMap(actorsArr []actors.Actor) map[string]actors.Actor {
	dataMap := make(map[string]actors.Actor)
	for _, actor := range actorsArr {
		dataMap[actor.ImdbID] = actor
	}
	return dataMap
}

func LoaderMiddleware(next http.Handler, actorService actors.Service, movieService movies.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ctx := NewLoaders(r.Context(), actorService, movieService)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
