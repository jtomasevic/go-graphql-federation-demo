package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/jtomasevic/go-graphql-federation-demo/actors/graph"
	"github.com/jtomasevic/go-graphql-federation-demo/actors/resolvers"
	"github.com/jtomasevic/go-graphql-federation-demo/actors/services"
	"github.com/rs/cors"
)

// NewGQLHttpHandler: factory method to return http handler for graphQL server.
func NewGQLHttpHandler() http.Handler {

	// get our services
	services := services.InilizeServices()
	// create graphQl server w
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &resolvers.Resolver{
				Services: *services,
			},
		}))
	// our UI will be run on localhost:3000, so we want to add cors rules to our
	// http handler.
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "*"},
		AllowCredentials: true,
	})
	// create http handler
	handler := c.Handler(srv)
	// load middleware for seave seas service data loaders
	// handler = seven_seas_dataloaders.LoaderMiddleware(handler, services.SevenSeasService)
	return handler
}
