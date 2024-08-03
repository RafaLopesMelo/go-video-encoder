package pg_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repo/pg"
	"github.com/RafaLopesMelo/go-video-encoder/test"
	"github.com/stretchr/testify/require"
)

func TestCreateVideo(t *testing.T) {
	connection := pg.NewConnection()
	repo := pg.NewVideosRepo(connection)

	dummy := test.DummyVideo()

	err := repo.Save(dummy)
	require.Nil(t, err)

	video, err := repo.FindByID(*dummy.Video().ID)

	require.Nil(t, err)
	require.Equal(t, dummy.Video().ID, video.ID)
}

func TestUpdateVideo(t *testing.T) {
	connection := pg.NewConnection()
	repo := pg.NewVideosRepo(connection)

	dummy := test.DummyVideo()
	repo.Save(dummy)

	video := dummy.Video()
	video.Status = entity.VideoStatusUploaded
	validated, _ := entity.NewValidatedVideo(video)

	repo.Save(validated)
	updated, err := repo.FindByID(*video.ID)

	require.Nil(t, err)
	require.Equal(t, video.ID, updated.ID)
	require.Equal(t, video.Status, updated.Status)
}
