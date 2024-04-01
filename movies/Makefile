PHONY: setup generate generate-dl

setup:
	# Add github.com/99designs/gqlgen to your project's tools.go
	printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
	go mod tidy

generate:
	go run github.com/99designs/gqlgen generate

run:
	go run server.go

generate-dl:
	cd services/movies; \
		go run github.com/vektah/dataloaden MoviesLoader string '[]github.com/jtomasevic/go-graphql-federation-demo/movies/services/movies.Movie'; \	
		# go run github.com/vektah/dataloaden ActorsLoader string 'github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors.Actor'; 