package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jtomasevic/go-graphql-federation-demo/actors/server"
)

const defaultPort = "4001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	// create http graphql handler
	handler := server.NewGQLHttpHandler()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// handle graphQL request
	http.Handle("/query", handler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
