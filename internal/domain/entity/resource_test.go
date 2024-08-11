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

	_, err := entity.NewValidatedResource(rv)

	require.NoError(t, err)
}

func TestRawVideoWithoutExtension(t *testing.T) {
	rv := entity.NewRawVideo(entity.NewRawVideoDto{
		Extension: "",
		NewResourceDto: entity.NewResourceDto{
			VideoID:         vo.NewID(),
			StorageProvider: entity.ResourceStorageProviderGCP,
			Path:            "/test",
			UploadURL:       "/test",
			Size:            100,
		},
	}, nil)

	rv.Resource().Status = entity.ResourceStatusPending
	_, err := entity.NewValidatedResource(rv)
	require.Nil(t, err)

	rv.Resource().Status = entity.ResourceStatusActive
	_, err = entity.NewValidatedResource(rv)
	require.ErrorIs(t, err, domainerrors.RequiredProperty)
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

	_, err := entity.NewValidatedResource(rv)

	require.ErrorIs(t, err, domainerrors.RequiredProperty)
}
