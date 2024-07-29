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
	ResourceStatusPending ResourceStatus = "PENDING"
	ResourceStatusActive  ResourceStatus = "ACTIVE"
	ResourceStatusDeleted ResourceStatus = "DELETED"
)

type resource struct {
	ID              *vo.UniqueEntityID
	Status          ResourceStatus
	Kind            ResourceKind
	VideoID         *vo.UniqueEntityID
	StorageProvider ResourceStorageProvider
	Path            string
	UploadURL       string
	Size            int
}

type ResourceWrapper interface {
	ID() vo.UniqueEntityID
	Resource() *resource
	validate() error
	Metadata() map[string]any
}

type NewResourceDto struct {
	VideoID         *vo.UniqueEntityID
	StorageProvider ResourceStorageProvider
	Path            string
	UploadURL       string
	Size            int
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

func NewResourceFromDto(dto LoadResourceDto, id *vo.UniqueEntityID) *resource {
	resource := resource{
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

func (resource *resource) validate() error {
	if resource.VideoID == nil {
		return domainerrors.RequiredProperty
	}

	return nil
}
