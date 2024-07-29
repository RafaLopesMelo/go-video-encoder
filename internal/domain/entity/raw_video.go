package entity

import (
	domainerrors "github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type RawVideo struct {
	resource  *resource
	Extension string
}

func (r RawVideo) ID() vo.UniqueEntityID {
	return *r.resource.ID
}

func (r RawVideo) Resource() *resource {
	return r.resource
}

func (r RawVideo) Metadata() map[string]any {
	return map[string]any{
		"extension": r.Extension,
	}
}

type NewRawVideoDto struct {
	NewResourceDto
	Extension string
}

func NewRawVideo(input NewRawVideoDto, id *vo.UniqueEntityID) *RawVideo {
	if id == nil {
		id = vo.NewID()
	}

	resource := &resource{
		ID:              id,
		Kind:            ResourceKindRawVideo,
		Status:          ResourceStatusPending,
		VideoID:         input.VideoID,
		StorageProvider: input.StorageProvider,
		Path:            input.Path,
		UploadURL:       input.UploadURL,
		Size:            input.Size,
	}

	return &RawVideo{
		resource:  resource,
		Extension: input.Extension,
	}
}

func (r RawVideo) validate() error {
	if r.Extension == "" {
		return domainerrors.RequiredProperty
	}

	return nil
}
