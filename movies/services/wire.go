// go:build wireinject
//go:build wireinject
// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/jtomasevic/go-graphql-federation-demo/movies/movies/services/movies"
	datasource "github.com/jtomasevic/go-graphql-federation-demo/movies/movies/services/movies/data_source"
	"github.com/jtomasevic/go-graphql-federation-demo/movies/movies/services/movies/data_store/repo"
)

func InilizeServices() *Services {
	panic(wire.Build(

		wire.Bind(new(datasource.DataSource), new(*datasource.DataSourceService)),
		datasource.NewDataSource,

		wire.Bind(new(movies.DataStore), new(*repo.MovieRepo)),
		movies.NewDataStore,

		wire.Bind(new(movies.Service), new(*movies.MovieService)),
		movies.NewService,

		wire.Struct(new(Services), "*"),
	))
}
