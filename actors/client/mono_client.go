package client

import "github.com/jtomasevic/go-graphql-federation-demo/actors/services"

func GetServices() services.Services {
	return *services.InilizeServices()
}
