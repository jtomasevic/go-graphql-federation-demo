package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"fmt"

	"github.com/jtomasevic/go-graphql-federation-demo/movies/graph"
	"github.com/jtomasevic/go-graphql-federation-demo/movies/graph/model"
	dataloaders "github.com/jtomasevic/go-graphql-federation-demo/movies/services/movies/data_loaders"
)

// FindActorByImdbID is the resolver for the findActorByImdbID field.
func (r *entityResolver) FindActorByImdbID(ctx context.Context, imdbID string) (model.Actor, error) {
	fmt.Printf("movie->entityResolver->FindActorByImdbID, imdbId: '%s'\n", imdbID)

	actor, err := dataloaders.GetLoaders(ctx).ActorsByImdbs.Load(imdbID)
	if err != nil {
		return model.Actor{}, err
	}
	return model.Actor{
		ID:           actor.ID,
		Name:         actor.Name,
		MovieImdbIds: actor.MovieImdbIds,
		ImdbID:       imdbID,
	}, nil

	// this is implementation without data resolvers, left here so we can learn/compare behaviour.
	// actor, err := r.ActorServices.ActorService.GetActorByImdbID(ctx, imdbID)
	// if err != nil {
	// 	return model.Actor{}, err
	// }
	// return model.Actor{
	// 	ID:           actor.ID,
	// 	Name:         actor.Name,
	// 	MovieImdbIds: actor.MovieImdbIds,
	// 	ImdbID:       imdbID,
	// }, nil
}

// Entity returns graph.EntityResolver implementation.
func (r *Resolver) Entity() graph.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
