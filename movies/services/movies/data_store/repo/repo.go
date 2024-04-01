package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	datasource "github.com/jtomasevic/go-graphql-federation-demo/movies/services/movies/data_source"
	"github.com/jtomasevic/go-graphql-federation-demo/movies/services/movies/data_store/model"
)

type MovieRepo struct {
	DataSource datasource.DataSource
}

func (repo *MovieRepo) GetMovies(ctx context.Context) ([]model.Movie, error) {
	query := "SELECT id, title, imdb_id FROM movie"
	fmt.Println(query)
	results := []model.Movie{}
	rows, err := repo.DataSource.Db().Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id uuid.UUID
		var title string
		var imdbId string
		err := rows.Scan(&id, &title, &imdbId)
		if err != nil {
			return nil, err
		}
		results = append(results, model.Movie{
			ID:     id,
			Title:  title,
			ImdbID: imdbId,
		})
	}
	return results, nil
}

func (repo *MovieRepo) GetMovieByImdbId(ctx context.Context, imdbId string) (model.Movie, error) {
	query := fmt.Sprintf("SELECT id, title, imdb_id FROM movie WHERE imdb_id = '%s", imdbId)
	fmt.Println(query)
	row := repo.DataSource.Db().QueryRow(query)

	var id uuid.UUID
	var title string
	var dbimdbId string
	err := row.Scan(&id, &title, &dbimdbId)
	if err != nil {
		return model.Movie{}, err
	}
	result := model.Movie{
		ID:     id,
		Title:  title,
		ImdbID: dbimdbId,
	}

	return result, nil
}

func (repo *MovieRepo) GetMoviesByImdbId(ctx context.Context, imdbIds []string) ([]model.Movie, error) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT id, title, imdb_id FROM movie WHERE imdb_id in (")
	for _, id := range imdbIds {
		queryBuilder.WriteString(fmt.Sprintf("'%s',", id))
	}
	query, _ := strings.CutSuffix(queryBuilder.String(), ",")
	query = fmt.Sprintf("%s)", query)
	fmt.Println(query)
	results := []model.Movie{}
	rows, err := repo.DataSource.Db().Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id uuid.UUID
		var title string
		var imdbId string
		err := rows.Scan(&id, &title, &imdbId)
		if err != nil {
			return nil, err
		}
		results = append(results, model.Movie{
			ID:     id,
			Title:  title,
			ImdbID: imdbId,
		})
	}
	return results, nil
}
