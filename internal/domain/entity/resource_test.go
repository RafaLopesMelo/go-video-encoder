package entity_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	domainerrors "github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
	"github.com/stretchr/testify/require"
)

func TestValidResource(t *testing.T) {
	resource := entity.NewResource(entity.NewResourceDto{
		Type:      entity.ResourceTypeRawVideo,
		VideoID:   vo.NewID(),
		Provider:  entity.ResourceStorageProviderGCP,
		Path:      "/test",
		UploadURL: "/test",
		Size:      100,
	}, nil)

	_, err := entity.NewValidatedResource(*resource)

	require.Nil(t, err)
}

func TestValidResourceWithoutVideoID(t *testing.T) {
	resource := entity.NewResource(entity.NewResourceDto{
		Type:      entity.ResourceTypeRawVideo,
		VideoID:   vo.NewID(),
		Provider:  entity.ResourceStorageProviderGCP,
		Path:      "/test",
		UploadURL: "/test",
		Size:      100,
	}, nil)

	_, err := entity.NewValidatedResource(*resource)

	require.ErrorIs(t, err, domainerrors.RequiredProperty)
}