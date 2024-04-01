package resolvers

import (
	actor_services "github.com/jtomasevic/go-graphql-federation-demo/actors/services"
	"github.com/jtomasevic/go-graphql-federation-demo/movies/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Services      services.Services
	ActorServices actor_services.Services
}

func checkError(errs []error) bool {
	if errs == nil {
		return false
	}
	for _, err := range errs {
		if err != nil {
			return true
		}
	}
	return false
}
