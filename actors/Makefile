PHONY: setup generate run

setup:
	# Add github.com/99designs/gqlgen to your project's tools.go
	printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
	go mod tidy

generate:
	go run github.com/99designs/gqlgen generate

install: 
	go get github.com/99designs/gqlgen
	go get github.com/proullon/ramsql
	go get github.com/google/uuid
	go get github.com/stretchr/testify/require
	go get github.com/vektah/dataloaden
	go get github.com/google/wire
	go get github.com/rs/cors

run:
	go run server.go