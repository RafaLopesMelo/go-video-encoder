package mysql_videos_repository_test

import (
	"fmt"
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/configs/env"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repositories/mysql"
	mysql_videos_repository "github.com/RafaLopesMelo/go-video-encoder/internal/infra/repositories/mysql/videos"
	"github.com/RafaLopesMelo/go-video-encoder/test"
	"github.com/stretchr/testify/require"
)

func TestSaveVideo(t *testing.T) {
    err := env.Load(".env")

    fmt.Println(err)

    connection := mysql.NewConnection()
    repo := mysql_videos_repository.NewMysqlVideosRepository(connection)

    dummy := test.DummyVideo()

    repo.Save(dummy)
    video, err:= repo.FindByID(*dummy.Video().ID)

    require.Nil(t, err)
    require.Equal(t, dummy.Video().ID, video.ID)
}
