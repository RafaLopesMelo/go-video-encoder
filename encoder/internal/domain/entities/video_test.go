package entities_test

import (
	"testing"
	"time"

	"github.com/RafaLopesMelo/go-video-encoder/domain/entities"
	"github.com/RafaLopesMelo/go-video-encoder/domain/value_objects"
	"github.com/stretchr/testify/require"
)

func TestValidVideo(t *testing.T) {
	video := entities.NewVideo()

	video.ID = *value_objects.NewID()
	video.ResourceID = "a"
	video.FilePath = "/teste"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}
