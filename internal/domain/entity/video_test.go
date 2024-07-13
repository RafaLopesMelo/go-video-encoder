package entity_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/stretchr/testify/require"
)

func TestValidVideo(t *testing.T) {
	video := entity.NewVideo(entity.NewVideoDto{
		ResourceID: "a",
		FilePath:   "/teste",
	}, nil)

	_, err := entity.NewValidatedVideo(*video)

	require.Nil(t, err)
}

func TestVideoWithoutFilePath(t *testing.T) {
	video := entity.NewVideo(entity.NewVideoDto{
		ResourceID: "a",
		FilePath:   "",
	}, nil)

	_, err := entity.NewValidatedVideo(*video)

	expected := errors.NewRequiredPropertyError("FilePath")

	require.EqualError(t, err, expected.Error())
}

func TestVideoWithoutResourceID(t *testing.T) {
	video := entity.NewVideo(entity.NewVideoDto{
		ResourceID: "",
		FilePath:   "/test",
	}, nil)

	_, err := entity.NewValidatedVideo(*video)

	expected := errors.NewRequiredPropertyError("ResourceID")

	require.EqualError(t, err, expected.Error())
}
