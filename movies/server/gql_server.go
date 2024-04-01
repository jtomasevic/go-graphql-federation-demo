package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/jtomasevic/go-graphql-federation-demo/actors/client"
	"github.com/jtomasevic/go-graphql-federation-demo/movies/graph"
	"github.com/jtomasevic/go-graphql-federation-demo/movies/resolvers"
	"github.com/jtomasevic/go-graphql-federation-demo/movies/services"
	dataloaders "github.com/jtomasevic/go-graphql-federation-demo/movies/services/movies/data_loaders"
	"github.com/jtomasevic/go-graphql-federation-demo/movies/services/movies/data_store/repo"
	"github.com/rs/cors"
)

// NewGQLHttpHandler: factory method to return http handler for graphQL server.
func NewGQLHttpHandler() http.Handler {
	// get our services
	repo.InitDb()
	repo.PopulateDb()
	actorClient := client.GetServices()
	services := services.InilizeServices()
	// create graphQl server w
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &resolvers.Resolver{
				Services:      *services,
				ActorServices: actorClient,
			},
			// Schema     *ast.Schema
			// Directives graph.DirectiveRoot,
			// Complexity graph.ComplexityRoot,
		}))
	// our UI will be run on localhost:3000, so we want to add cors rules to our
	// http handler.
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "*"},
		AllowCredentials: true,
	})
	// create http handler
	handler := c.Handler(srv)
	// load middleware for data loaders
	handler = dataloaders.LoaderMiddleware(handler, actorClient.ActorService, services.MovieService)
	return handler
}
