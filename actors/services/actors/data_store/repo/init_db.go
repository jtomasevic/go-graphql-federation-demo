package repo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors/data_store/model"
	_ "github.com/proullon/ramsql/driver"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("ramsql", "Actors")
	if err != nil {
		log.Fatalf("sql.Open : Error : %s\n", err)
		panic(err)
	}

	return db
}

func InitDb() error {
	batch := []string{
		`CREATE TABLE actor (id UUID PRIMARY KEY, name TEXT, movie_imdbs TEXT, imdb_id TEXT);`,
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
	actors := createTestObjects()
	batch := []string{}

	for _, actor := range actors {
		batch = append(batch,
			fmt.Sprintf("INSERT INTO actor (id, name, movie_imdbs, imdb_id) VALUES ('%s', '%s', '%s', '%s');",
				actor.ID.String(), actor.Name, actor.MovieImdbs, actor.ImdbID))
	}

	db := GetConnection()
	defer db.Close()

	for _, b := range batch {
		// fmt.Println(b)
		_, err := db.Exec(b)
		if err != nil {
			log.Fatalf("sql.Exec: Error: %s\n", err)
			return err
		}
	}

	return nil
}

func createTestObjects() []model.Actor {
	return []model.Actor{
		{
			ID:         uuid.MustParse("4a1230c5-2367-4b98-841e-e735edfb263e"),
			Name:       "Johnny Depp",
			MovieImdbs: "tt4123430;tt0325980;tt1152836",
			ImdbID:     "nm0000136",
		},
		{
			ID:         uuid.MustParse("356bcd5b-cf8d-4062-9c85-8267b4be5b1c"),
			Name:       "Geoffrey Rush",
			MovieImdbs: "tt0138097;tt1504320;tt0180073",
			ImdbID:     "nm0001691",
		},
		{
			ID:         uuid.MustParse("a3630f0a-3b0c-499c-b962-28c9ded58df8"),
			Name:       "Cillian Murphy",
			MovieImdbs: "tt0289043;tt2442560;tt14569898",
			ImdbID:     "nm0614165",
		},
		{
			ID:         uuid.MustParse("22a8eca6-4119-4eab-89a4-2f61d657c578"),
			Name:       "Leonardo DiCaprio",
			MovieImdbs: "tt1375666;tt0407887;tt0217505",
			ImdbID:     "nm0000138",
		},
		{
			ID:         uuid.MustParse("970b83cd-7240-4ae1-b20c-db1aeb20a02f"),
			Name:       "Brad Pitt",
			MovieImdbs: "tt0114369;tt0137523;tt0114746",
			ImdbID:     "nm0000093",
		},
	}
}
