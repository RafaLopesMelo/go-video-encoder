package entity_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	domainerrors "github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
	"github.com/stretchr/testify/require"
)

func TestValidRawVideo(t *testing.T) {
	rv := entity.NewRawVideo(entity.NewRawVideoDto{
		Extension: "mp4",
		NewResourceDto: entity.NewResourceDto{
			VideoID:         vo.NewID(),
			StorageProvider: entity.ResourceStorageProviderGCP,
			Path:            "/test",
			UploadURL:       "/test",
			Size:            100,
		},
	}, nil)

	_, err := entity.NewValidatedResource(*rv)

	require.Nil(t, err)
}

func TestResourceWithoutVideoID(t *testing.T) {
	rv := entity.NewRawVideo(entity.NewRawVideoDto{
		Extension: "mp4",
		NewResourceDto: entity.NewResourceDto{
			VideoID:         nil,
			StorageProvider: entity.ResourceStorageProviderGCP,
			Path:            "/test",
			UploadURL:       "/test",
			Size:            100,
		},
	}, nil)

	_, err := entity.NewValidatedResource(*rv)

	require.ErrorIs(t, err, domainerrors.RequiredProperty)
}
