package entity

import (
	domainerrors "github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type ResourceType string

const (
	ResourceTypeRawVideo        ResourceType = "RAW_VIDEO"
	ResourceTypeTranscodedVideo ResourceType = "TRANSCODED_VIDEO"
)

type Resource struct {
	ID      *vo.UniqueEntityID
	Type    ResourceType
	VideoID *vo.UniqueEntityID
	Storage *vo.Storage
}

type NewResourceDto struct {
	Type    ResourceType
	VideoID *vo.UniqueEntityID
	Storage *vo.Storage
}

func NewResource(input NewResourceDto, id *vo.UniqueEntityID) *Resource {
	if id == nil {
		id = vo.NewID()
	}

	resource := Resource{
		ID:      id,
		Type:    input.Type,
		VideoID: input.VideoID,
		Storage: input.Storage,
	}

	return &resource
}

func (resource *Resource) validate() error {
	if resource.VideoID == nil {
		return domainerrors.RequiredProperty
	}

	if resource.Storage == nil {
		return domainerrors.RequiredProperty
	}

	return nil
}
