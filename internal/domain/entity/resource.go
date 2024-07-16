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

type ResourceStorageProvider string

const (
	ResourceStorageProviderGCP ResourceStorageProvider = "GCP"
)

type ResourceStatus string

const (
	ResourceStatusActive  ResourceStatus = "ACTIVE"
	ResourceStatusDeleted ResourceStatus = "DELETED"
)

type Resource struct {
	ID        *vo.UniqueEntityID
	Status    ResourceStatus
	Type      ResourceType
	VideoID   *vo.UniqueEntityID
	Provider  ResourceStorageProvider
	Path      string
	UploadURL string
	Size      int
}

type NewResourceDto struct {
	Type      ResourceType
	VideoID   *vo.UniqueEntityID
	Provider  ResourceStorageProvider
	Path      string
	UploadURL string
	Size      int
}

func NewResource(input NewResourceDto, id *vo.UniqueEntityID) *Resource {
	if id == nil {
		id = vo.NewID()
	}

	resource := Resource{
		ID:        id,
		Type:      input.Type,
		Status:    ResourceStatusActive,
		VideoID:   input.VideoID,
		Provider:  input.Provider,
		Path:      input.Path,
		UploadURL: input.UploadURL,
		Size:      input.Size,
	}

	return &resource
}

type LoadResourceDto struct {
	Status    ResourceStatus
	Type      ResourceType
	VideoID   *vo.UniqueEntityID
	Provider  ResourceStorageProvider
	Path      string
	UploadURL string
	Size      int
}

func NewResourceFromDto(dto LoadResourceDto, id *vo.UniqueEntityID) *Resource {
	resource := Resource{
		ID:        id,
		Status:    dto.Status,
		Type:      dto.Type,
		VideoID:   dto.VideoID,
		Provider:  dto.Provider,
		Path:      dto.Path,
		UploadURL: dto.UploadURL,
		Size:      dto.Size,
	}

	return &resource
}

func (resource *Resource) validate() error {
	if resource.VideoID == nil {
		return domainerrors.RequiredProperty
	}

	return nil
}
