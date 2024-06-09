package videos_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/videos"
	"github.com/stretchr/testify/require"
)

func TestValidVideo(t *testing.T) {
	_, err := videos.NewVideo(videos.NewVideoDto{
        ResourceID: "a",
        FilePath: "/teste",
    }, nil)

	require.Nil(t, err)
}
