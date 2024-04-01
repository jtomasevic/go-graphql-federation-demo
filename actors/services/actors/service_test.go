package actors

import (
	"context"
	"fmt"
	"testing"

	datasource "github.com/jtomasevic/go-graphql-federation-demo/actors/services/actors/data_source"
	"github.com/stretchr/testify/require"
)

func TestGetByAttribute(t *testing.T) {
	service := NewService(NewDataStore(datasource.NewDataSource()))

	t.Run("GetByImdbId", func(t *testing.T) {
		// // check for valid id in init_db.go/createTestObjects
		actorID := "4a1230c5-2367-4b98-841e-e735edfb263e"
		actor, err := service.GetActorByImdbID(context.TODO(), "nm0000136")
		require.NoError(t, err)
		require.Equal(t, actorID, actor.ID.String())
		require.Equal(t, "Johnny Depp", actor.Name)
		require.Equal(t, []string{"tt4123430", "tt0325980", "tt1152836"}, actor.MovieImdbIds)
		require.Equal(t, "nm0000136", actor.ImdbID)
	})

	t.Run("GetActorsByImdbIDs", func(t *testing.T) {
		// // check for valid ids in init_db.go/createTestObjects
		actorIDs := []string{"4a1230c5-2367-4b98-841e-e735edfb263e", "356bcd5b-cf8d-4062-9c85-8267b4be5b1c"}
		actors, err := service.GetActorsByImdbIDs(context.TODO(), []string{"nm0000136", "nm0001691"})

		require.NoError(t, err)
		require.Len(t, actors, 2)
		require.Equal(t, actorIDs[0], actors[0].ID.String())
		require.Equal(t, "Johnny Depp", actors[0].Name)
		require.Equal(t, []string{"tt4123430", "tt0325980", "tt1152836"}, actors[0].MovieImdbIds)
		require.Equal(t, "nm0000136", actors[0].ImdbID)
		require.Equal(t, actorIDs[1], actors[1].ID.String())
		require.Equal(t, "Geoffrey Rush", actors[1].Name)
		require.Equal(t, []string{"tt0138097", "tt1504320", "tt0180073"}, actors[1].MovieImdbIds)
		require.Equal(t, "nm0001691", actors[1].ImdbID)
		fmt.Println("hej")
	})
}
