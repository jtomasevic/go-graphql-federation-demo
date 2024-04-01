package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	datasource "github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors/data_source"
	"github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors/data_store/model"
)

type ActorRepo struct {
	DataSource datasource.DataSource
}

func (repo *ActorRepo) GetActors(ctx context.Context) ([]model.Actor, error) {
	query := "select id, name, movie_imdbs, imdb_id from actor"
	fmt.Println(query)
	results := []model.Actor{}
	rows, err := repo.DataSource.Db().Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id uuid.UUID
		var name string
		var movieImdbs string
		var imdbId string

		err := rows.Scan(&id, &name, &movieImdbs, &imdbId)
		if err != nil {
			return nil, err
		}
		results = append(results, model.Actor{
			ID:         id,
			Name:       name,
			MovieImdbs: movieImdbs,
			ImdbID:     imdbId,
		})
	}
	return results, nil
}

// func (repo *ActorRepo) GetActorById(ctx context.Context, id uuid.UUID) (model.Actor, error) {
// 	query := fmt.Sprintf("select id, name, movie_imdbs, imdb_id from actor where id = '%s'", id.String())
// 	fmt.Println(query)
// 	row:= repo.DataSource.Db().QueryRow(query)
// 	var dbid uuid.UUID
// 	var name string
// 	var movieImdbs string
// 	var imdbId string
// 	err := row.Scan(&dbid, &name,  &movieImdbs,  &imdbId)
// 	if err != nil {
// 		return model.Actor{}, err
// 	}
// 	result := model.Actor{
// 		ID:   dbid,
// 		Name: name,
// 		MovieImdbs: movieImdbs,
// 		ImdbID: imdbId,
// 	}
// 	return result, nil
// }

func (repo *ActorRepo) GetActorByName(ctx context.Context, name string) (model.Actor, error) {
	query := fmt.Sprintf("select id, name, movie_imdbs, imdb_id from actor where name = '%s'", name)
	fmt.Println(query)
	row := repo.DataSource.Db().QueryRow(query)
	var dbid uuid.UUID
	var dbname string
	var movieImdbs string
	var imdbId string
	err := row.Scan(&dbid, &dbname, &movieImdbs, &imdbId)
	if err != nil {
		return model.Actor{}, err
	}
	result := model.Actor{
		ID:         dbid,
		Name:       dbname,
		MovieImdbs: movieImdbs,
		ImdbID:     imdbId,
	}
	return result, nil
}

func (repo *ActorRepo) GetActorByImdbID(ctx context.Context, imdbId string) (model.Actor, error) {
	query := fmt.Sprintf("select id, name, movie_imdbs, imdb_id from actor where imdb_id = '%s'", imdbId)
	fmt.Println(query)
	row := repo.DataSource.Db().QueryRow(query)
	var dbid uuid.UUID
	var dbname string
	var movieImdbs string
	var dbimdbId string
	err := row.Scan(&dbid, &dbname, &movieImdbs, &dbimdbId)
	if err != nil {
		return model.Actor{}, err
	}
	result := model.Actor{
		ID:         dbid,
		Name:       dbname,
		MovieImdbs: movieImdbs,
		ImdbID:     dbimdbId,
	}
	return result, nil
}

func (repo *ActorRepo) GetActorsByImdbIDs(ctx context.Context, imdbId []string) ([]model.Actor, error) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString("select id, name, movie_imdbs, imdb_id from actor where imdb_id IN (")
	for _, id := range imdbId {
		queryBuilder.WriteString(fmt.Sprintf("'%s',", id))
	}
	query, _ := strings.CutSuffix(queryBuilder.String(), ",")
	query = fmt.Sprintf("%s)", query)
	fmt.Println(query)

	results := []model.Actor{}
	rows, err := repo.DataSource.Db().Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id uuid.UUID
		var name string
		var movieImdbs string
		var imdbId string

		err := rows.Scan(&id, &name, &movieImdbs, &imdbId)
		if err != nil {
			return nil, err
		}
		results = append(results, model.Actor{
			ID:         id,
			Name:       name,
			MovieImdbs: movieImdbs,
			ImdbID:     imdbId,
		})
	}
	return results, nil
}
