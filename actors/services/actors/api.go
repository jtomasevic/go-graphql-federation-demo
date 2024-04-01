package actors

import (
	"context"

	datasource "github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors/data_source"
	"github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors/data_store/model"
	"github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors/data_store/repo"
)

type Service interface {
	GetActors(ctx context.Context) ([]Actor, error)
	// GetActorById(ctx context.Context, id uuid.UUID) (Actor, error)
	GetActorByName(ctx context.Context, name string) (Actor, error)
	GetActorByImdbID(ctx context.Context, imdbId string) (Actor, error)
	GetActorsByImdbIDs(ctx context.Context, imdbId []string) ([]Actor, error)
}

type DataStore interface {
	GetActors(ctx context.Context) ([]model.Actor, error)
	// GetActorById(ctx context.Context, id uuid.UUID) (model.Actor, error)
	GetActorByName(ctx context.Context, name string) (model.Actor, error)
	GetActorByImdbID(ctx context.Context, imdbId string) (model.Actor, error)
	GetActorsByImdbIDs(ctx context.Context, imdbId []string) ([]model.Actor, error)
}

func NewDataStore(dataSource datasource.DataSource) *repo.ActorRepo {
	repo.InitDb()
	repo.PopulateDb()
	return &repo.ActorRepo{
		DataSource: dataSource,
	}
}

func NewService(dataStore DataStore) *ActorService {
	return &ActorService{
		dataStore: dataStore,
	}
}
