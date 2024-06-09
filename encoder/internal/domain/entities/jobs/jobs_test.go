package jobs_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/jobs"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/videos"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
    video, _ := videos.NewVideo(videos.NewVideoDto{
        FilePath: "/teste",
        ResourceID: "a",
    }, nil)

    job, err := jobs.NewJob("/path", "CONVERTED", video)

    require.Nil(t, err)
    require.NotNil(t, job)
}
