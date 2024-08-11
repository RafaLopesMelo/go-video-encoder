package video_test

import (
	"os"
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/app/usecase/video"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/config/env"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repo/pg"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	err := env.Load(".env")

	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestRegisterVideo(t *testing.T) {
	connection := pg.NewConnection()
	rv := pg.NewVideosRepo(connection)
	rr := pg.NewResourcesRepo(connection)

	uc := video.NewRegisterUseCase(rv, rr)

	created, err := uc.Execute()

	require.NoError(t, err)
	require.NotEmpty(t, created.ID)
	require.NotEmpty(t, created.UploadURL)
}
