package pg_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repository/pg"
	"github.com/RafaLopesMelo/go-video-encoder/test"
	"github.com/stretchr/testify/require"
)

func TestCreateVideo(t *testing.T) {
	connection := pg.NewConnection()
	repo := pg.NewVideosRepository(connection)

	dummy := test.DummyVideo()

	repo.Save(dummy)
	video, err := repo.FindByID(*dummy.Video().ID)

	require.Nil(t, err)
	require.Equal(t, dummy.Video().ID, video.ID)
}

func TestUpdateVideo(t *testing.T) {
	connection := pg.NewConnection()
	repo := pg.NewVideosRepository(connection)

	dummy := test.DummyVideo()
	repo.Save(dummy)

	video := dummy.Video()
	video.FilePath = "/test-updated"
	validated, _ := entity.NewValidatedVideo(video)

	repo.Save(validated)
	updated, err := repo.FindByID(*video.ID)

	require.Nil(t, err)
	require.Equal(t, video.ID, updated.ID)
	require.Equal(t, video.FilePath, updated.FilePath)
}
