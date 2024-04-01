package repo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jtomasevic/go-graphql-federation-demo/movies/services/movies/data_store/model"
	_ "github.com/proullon/ramsql/driver"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("ramsql", "Movies")
	if err != nil {
		log.Fatalf("sql.Open : Error : %s\n", err)
		panic(err)
	}

	return db
}

func InitDb() error {
	batch := []string{
		`CREATE TABLE movie (id UUID PRIMARY KEY, title TEXT, imdb_id TEXT);`,
	}
	db := GetConnection()

	defer db.Close()

	for _, b := range batch {
		_, err := db.Exec(b)
		if err != nil {
			log.Fatalf("sql.Exec: Error: %s\n", err)
			return err
		}
	}
	return nil
}

func PopulateDb() error {
	movies := createTestObjects()
	batch := []string{}

	for _, movie := range movies {
		batch = append(batch,
			fmt.Sprintf("INSERT INTO movie (id, title, imdb_id) VALUES ('%s', '%s', '%s');",
				movie.ID.String(), movie.Title, movie.ImdbID))
	}

	db := GetConnection()
	defer db.Close()

	for _, b := range batch {
		_, err := db.Exec(b)
		if err != nil {
			log.Fatalf("sql.Exec: Error: %s\n", err)
			return err
		}
	}
	return nil
}

func createTestObjects() []model.Movie {
	return []model.Movie{
		{
			ID:     uuid.New(),
			Title:  "Fantastic Beasts: The Crimes of Grindelwald",
			ImdbID: "tt4123430",
		},
		{
			ID:     uuid.New(),
			Title:  "Pirates of the Caribbean: The Curse of the Black Pearl",
			ImdbID: "tt0325980",
		},
		{
			ID:     uuid.New(),
			Title:  "Public Enemies",
			ImdbID: "tt1152836",
		},
		{
			ID:     uuid.New(),
			Title:  "Shakespeare in Love",
			ImdbID: "tt0138097",
		},
		{
			ID:     uuid.New(),
			Title:  "The Kings Speech",
			ImdbID: "tt1504320",
		},
		{
			ID:     uuid.New(),
			Title:  "Quills",
			ImdbID: "tt0180073",
		},
		{
			ID:     uuid.New(),
			Title:  "28 Days Later...",
			ImdbID: "tt0289043",
		},
		{
			ID:     uuid.New(),
			Title:  "Peaky Blinders",
			ImdbID: "tt2442560",
		},
		{
			ID:     uuid.New(),
			Title:  "All of This Unreal Time",
			ImdbID: "tt14569898",
		},
		{
			ID:     uuid.New(),
			Title:  "Inception",
			ImdbID: "tt1375666",
		},
		{
			ID:     uuid.New(),
			Title:  "The Departed",
			ImdbID: "tt0407887",
		},
		{
			ID:     uuid.New(),
			Title:  "Gangs of New York",
			ImdbID: "tt0217505",
		},
		{
			ID:     uuid.New(),
			Title:  "Seven",
			ImdbID: "tt0114369",
		},
		{
			ID:     uuid.New(),
			Title:  "Fight Club",
			ImdbID: "tt0137523",
		},
		{
			ID:     uuid.New(),
			Title:  "Twelve Monkeys",
			ImdbID: "tt0114746",
		},
	}
}
