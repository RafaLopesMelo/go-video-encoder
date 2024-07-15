package vo

import domainerrors "github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"

type StorageProvider string

const (
	StorageProviderGCP StorageProvider = "GCP"
)

type Storage struct {
	Provider  StorageProvider
	Path      string
	UploadURL string
}

type NewStorageDto struct {
	Provider  StorageProvider
	Path      string
	UploadURL string
}

func NewStorage(input NewStorageDto) (*Storage, error) {
	if input.Provider == "" {
		return nil, domainerrors.RequiredProperty
	}

	if input.Path == "" {
		return nil, domainerrors.RequiredProperty
	}

	storage := Storage{
		Provider:  input.Provider,
		Path:      input.Path,
		UploadURL: input.UploadURL,
	}

	return &storage, nil
}
