package actors

import (
	"context"

)

type ActorService struct {
	dataStore DataStore
}

func(service *ActorService) GetActors(ctx context.Context) ([]Actor, error) {
	actors, err:= service.dataStore.GetActors(ctx)
	if err != nil {
		return nil, err
	}
	return dbToServiceActors(actors), nil
}

// func(service *ActorService) GetActorById(ctx context.Context, id uuid.UUID) (Actor, error) {
// 	actor, err:= service.dataStore.GetActorById(ctx, id)
// 	if err != nil {
// 		return Actor{}, err
// 	}
// 	return dbToServiceActor(actor), nil
// }

func(service *ActorService) GetActorByName(ctx context.Context, name string) (Actor, error) {
	actor, err:= service.dataStore.GetActorByName(ctx, name)
	if err != nil {
		return Actor{}, err
	}
	return dbToServiceActor(actor), nil
}

func (service *ActorService) GetActorByImdbID(ctx context.Context, imdbId string) (Actor, error) {
	actor, err:= service.dataStore.GetActorByImdbID(ctx, imdbId)
	if err != nil {
		return Actor{}, err
	}
	return dbToServiceActor(actor), nil
}

func (service *ActorService) GetActorsByImdbIDs(ctx context.Context, imdbId []string) ([]Actor, error) {
	actors, err:= service.dataStore.GetActorsByImdbIDs(ctx, imdbId)
	if err != nil {
		return nil, err
	}
	return dbToServiceActors(actors), nil
}