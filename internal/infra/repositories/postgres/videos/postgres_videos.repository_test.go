package postgres_videos_repository_test

import (
	"os"
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/configs/env"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repositories/postgres"
	postgres_videos_repository "github.com/RafaLopesMelo/go-video-encoder/internal/infra/repositories/postgres/videos"
	"github.com/RafaLopesMelo/go-video-encoder/test"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	err := env.Load(".env")

	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestCreateVideo(t *testing.T) {
	connection := postgres.NewConnection()
	repo := postgres_videos_repository.NewPostgresVideosRepository(connection)

	dummy := test.DummyVideo()

	repo.Save(dummy)
	video, err := repo.FindByID(*dummy.Video().ID)

	require.Nil(t, err)
	require.Equal(t, dummy.Video().ID, video.ID)
}

func TestUpdateVideo(t *testing.T) {
	connection := postgres.NewConnection()
	repo := postgres_videos_repository.NewPostgresVideosRepository(connection)

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
