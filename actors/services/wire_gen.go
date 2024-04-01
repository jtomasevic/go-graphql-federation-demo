// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package services

import (
	"github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors"
	"github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors/data_source"
)

// Injectors from wire.go:

func InilizeServices() *Services {
	dataSourceService := datasource.NewDataSource()
	actorRepo := actors.NewDataStore(dataSourceService)
	actorService := actors.NewService(actorRepo)
	services := &Services{
		ActorService: actorService,
	}
	return services
}

