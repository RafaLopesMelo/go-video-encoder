package entity_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	domainerrors "github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
	"github.com/RafaLopesMelo/go-video-encoder/test"
	"github.com/stretchr/testify/require"
)

func TestValidResource(t *testing.T) {
	resource := entity.NewResource(entity.NewResourceDto{
		Type:    entity.ResourceTypeRawVideo,
		VideoID: vo.NewID(),
		Storage: test.DummyStorage(),
	}, nil)

	_, err := entity.NewValidatedResource(*resource)

	require.Nil(t, err)
}

func TestValidResourceWithoutVideoID(t *testing.T) {
	resource := entity.NewResource(entity.NewResourceDto{
		Type:    entity.ResourceTypeRawVideo,
		VideoID: nil,
		Storage: test.DummyStorage(),
	}, nil)

	_, err := entity.NewValidatedResource(*resource)

	require.ErrorIs(t, err, domainerrors.RequiredProperty)
}

func TestValidResourceWithoutStorage(t *testing.T) {
	resource := entity.NewResource(entity.NewResourceDto{
		Type:    entity.ResourceTypeRawVideo,
		VideoID: vo.NewID(),
		Storage: nil,
	}, nil)

	_, err := entity.NewValidatedResource(*resource)

	require.ErrorIs(t, err, domainerrors.RequiredProperty)
}
