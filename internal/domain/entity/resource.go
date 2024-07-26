package entity

import (
	domainerrors "github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type ResourceKind string

const (
	ResourceKindRawVideo        ResourceKind = "RAW_VIDEO"
	ResourceKindTranscodedVideo ResourceKind = "TRANSCODED_VIDEO"
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
	ID              *vo.UniqueEntityID
	Status          ResourceStatus
	Kind            ResourceKind
	VideoID         *vo.UniqueEntityID
	StorageProvider ResourceStorageProvider
	Path            string
	UploadURL       string
	Size            int
}

type NewResourceDto struct {
	Kind            ResourceKind
	VideoID         *vo.UniqueEntityID
	StorageProvider ResourceStorageProvider
	Path            string
	UploadURL       string
	Size            int
}

func NewResource(input NewResourceDto, id *vo.UniqueEntityID) *Resource {
	if id == nil {
		id = vo.NewID()
	}

	resource := Resource{
		ID:              id,
		Kind:            input.Kind,
		Status:          ResourceStatusActive,
		VideoID:         input.VideoID,
		StorageProvider: input.StorageProvider,
		Path:            input.Path,
		UploadURL:       input.UploadURL,
		Size:            input.Size,
	}

	return &resource
}

type LoadResourceDto struct {
	Status          ResourceStatus
	Kind            ResourceKind
	VideoID         *vo.UniqueEntityID
	StorageProvider ResourceStorageProvider
	Path            string
	UploadURL       string
	Size            int
}

func NewResourceFromDto(dto LoadResourceDto, id *vo.UniqueEntityID) *Resource {
	resource := Resource{
		ID:              id,
		Status:          dto.Status,
		Kind:            dto.Kind,
		VideoID:         dto.VideoID,
		StorageProvider: dto.StorageProvider,
		Path:            dto.Path,
		UploadURL:       dto.UploadURL,
		Size:            dto.Size,
	}

	return &resource
}

func (resource *Resource) validate() error {
	if resource.VideoID == nil {
		return domainerrors.RequiredProperty
	}

	return nil
}
