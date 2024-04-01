// go:build wireinject
//go:build wireinject
// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors"
	datasource "github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors/data_source"
	"github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors/data_store/repo"
)

func InilizeServices() *Services {
	panic(wire.Build(

		wire.Bind(new(datasource.DataSource), new(*datasource.DataSourceService)),
		datasource.NewDataSource,

		wire.Bind(new(actors.DataStore), new(*repo.ActorRepo)),
		actors.NewDataStore,

		wire.Bind(new(actors.Service), new(*actors.ActorService)),
		actors.NewService,

		wire.Struct(new(Services), "*"),
	))
}
