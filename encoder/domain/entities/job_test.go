package entities_test

import (
	"testing"
	"time"

	"github.com/RafaLopesMelo/go-video-encoder/domain/entities"
	"github.com/RafaLopesMelo/go-video-encoder/domain/value_objects"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
    video := entities.NewVideo()
    video.ID = *value_objects.NewID()
    video.FilePath = "/test"
    video.CreatedAt = time.Now()

    job, err := entities.NewJob("/path", "CONVERTED", video)

    require.Nil(t, err)
    require.NotNil(t, job)
}
