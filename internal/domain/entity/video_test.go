package entity_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/stretchr/testify/require"
)

func TestValidVideo(t *testing.T) {
	video := entity.NewVideo(entity.NewVideoDto{}, nil)

	_, err := entity.NewValidatedVideo(*video)

	require.Nil(t, err)
}
