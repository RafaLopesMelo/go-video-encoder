package pg_test

import (
	"os"
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/config/env"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repo/pg"
	"github.com/RafaLopesMelo/go-video-encoder/test"
)

var videoId *vo.UniqueEntityID

func TestMain(m *testing.M) {
	err := env.Load(".env")

	if err != nil {
		panic(err)
	}

	connection := pg.NewConnection()
	repo := pg.NewVideosRepo(connection)

	video := test.DummyVideo()
	repo.Save(video)
	videoId = video.Video().ID

	os.Exit(m.Run())
}
