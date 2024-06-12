package mysql_videos_repository_test

import (
	"os"
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/videos"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/configs/env"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repositories/mysql"
	mysql_videos_repository "github.com/RafaLopesMelo/go-video-encoder/internal/infra/repositories/mysql/videos"
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
	connection := mysql.NewConnection()
	repo := mysql_videos_repository.NewMysqlVideosRepository(connection)

	dummy := test.DummyVideo()

	repo.Save(dummy)
	video, err := repo.FindByID(*dummy.Video().ID)

	require.Nil(t, err)
	require.Equal(t, dummy.Video().ID, video.ID)
}

func TestUpdateVideo(t *testing.T) {
	connection := mysql.NewConnection()
	repo := mysql_videos_repository.NewMysqlVideosRepository(connection)

	dummy := test.DummyVideo()
	repo.Save(dummy)

	video := dummy.Video()
	video.FilePath = "/test-updated"
	validated, _ := videos.NewValidatedVideo(video)

	repo.Save(validated)
	updated, err := repo.FindByID(*video.ID)

	require.Nil(t, err)
	require.Equal(t, video.ID, updated.ID)
	require.Equal(t, video.FilePath, updated.FilePath)
}
