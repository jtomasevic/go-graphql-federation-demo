package repo

import (
	"context"
	"testing"

	datasource "github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors/data_source"
	"github.com/stretchr/testify/require"
)

func TestGetByAttribute(t *testing.T) {
	InitDb()
	PopulateDb()
	repo := ActorRepo{
		DataSource: datasource.NewDataSource(),
	}
	// t.Run("GetActorById", func(t *testing.T) {
	// 	// check for valid id in init_db.go/createTestObjects
	// 	actor, err := repo.GetActorById(context.TODO(), uuid.MustParse("4a1230c5-2367-4b98-841e-e735edfb263e"))
	// 	require.NoError(t, err)
	// 	require.Equal(t, "4a1230c5-2367-4b98-841e-e735edfb263e", actor.ID.String())
	// 	require.Equal(t, "Johnny Depp", actor.Name)
	// 	require.Equal(t, "tt4123430;tt0325980;tt1152836", actor.MovieImdbs)
	// 	require.Equal(t, "nm0000136", actor.ImdbID)
	// })

	t.Run("GetActorByName", func(t *testing.T) {
		// check for valid names in init_db.go/createTestObjects
		actor, err := repo.GetActorByName(context.TODO(), "Johnny Depp")
		require.NoError(t, err)
		require.Equal(t, "4a1230c5-2367-4b98-841e-e735edfb263e", actor.ID.String())
		require.Equal(t, "Johnny Depp", actor.Name)
		require.Equal(t, "tt4123430;tt0325980;tt1152836", actor.MovieImdbs)
		require.Equal(t, "nm0000136", actor.ImdbID)
	})

	t.Run("GetActorByImdbID", func(t *testing.T) {
		// check for valid imdbId in init_db.go/createTestObjects
		actor, err := repo.GetActorByImdbID(context.TODO(), "nm0000136")
		require.NoError(t, err)
		require.Equal(t, "4a1230c5-2367-4b98-841e-e735edfb263e", actor.ID.String())
		require.Equal(t, "Johnny Depp", actor.Name)
		require.Equal(t, "tt4123430;tt0325980;tt1152836", actor.MovieImdbs)
		require.Equal(t, "nm0000136", actor.ImdbID)
	})

}

func TestGetByListOfAttributes(t *testing.T) {

	InitDb()
	PopulateDb()
	repo := ActorRepo{
		DataSource: datasource.NewDataSource(),
	}
	t.Run("GetActorByImdbs", func(t *testing.T) {
		actor, err := repo.GetActorsByImdbIDs(context.TODO(), []string{"nm0000136", "nm0001691", "nm0614165", "nm0000138", "nm0000093"})
		require.NoError(t, err)

		// check for valid imdbId in init_db.go/createTestObjects
		require.Equal(t, "Johnny Depp", actor[0].Name)
		require.Equal(t, "Geoffrey Rush", actor[1].Name)
		require.Equal(t, "Cillian Murphy", actor[2].Name)
		require.Equal(t, "Leonardo DiCaprio", actor[3].Name)
		require.Equal(t, "Brad Pitt", actor[4].Name)

	})
}
