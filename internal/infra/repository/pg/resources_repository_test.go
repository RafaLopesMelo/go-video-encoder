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

	dummy := test.DummyRawVideo(videoId)
	rw := dummy.Wrapper()

	err := repo.Save(dummy)
	require.Nil(t, err)

	resource, err := repo.FindByID(rw.ID())

	require.Nil(t, err)
	require.Equal(t, rw.ID(), resource.ID())
}

func TestUpdateResource(t *testing.T) {
	connection := pg.NewConnection()
	repo := pg.NewResourcesRepository(connection)

	dummy := test.DummyRawVideo(videoId)
	repo.Save(dummy)

	rw := dummy.Wrapper()
	resource := rw.Resource()
	resource.Path = "/test"
	validated, _ := entity.NewValidatedResource(rw)

	repo.Save(validated)
	updated, err := repo.FindByID(rw.ID())

	require.Nil(t, err)
	require.Equal(t, rw.ID(), updated.ID())
	require.Equal(t, resource.Path, updated.Resource().Path)
}
