package pg_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repository/pg"
	"github.com/RafaLopesMelo/go-video-encoder/test"
	"github.com/stretchr/testify/require"
)

func TestCreateResource(t *testing.T) {
	connection := pg.NewConnection()
	repo := pg.NewResourcesRepository(connection)

	dummy := test.DummyResource(videoId)

	err := repo.Save(dummy)
	require.Nil(t, err)

	resource, err := repo.FindByID(*dummy.Resource().ID)

	require.Nil(t, err)
	require.Equal(t, dummy.Resource().ID, resource.ID)
}

func TestUpdateResource(t *testing.T) {
	connection := pg.NewConnection()
	repo := pg.NewResourcesRepository(connection)

	dummy := test.DummyResource(videoId)
	repo.Save(dummy)

	resource := dummy.Resource()
	resource.Path = "/test"
	validated, _ := entity.NewValidatedResource(resource)

	repo.Save(validated)
	updated, err := repo.FindByID(*resource.ID)

	require.Nil(t, err)
	require.Equal(t, resource.ID, updated.ID)
	require.Equal(t, resource.Path, updated.Path)
}
