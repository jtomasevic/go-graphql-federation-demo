package resolvers

import "github.com/jtomasevic/go-graphql-federation-demo/actors/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Services services.Services
}
