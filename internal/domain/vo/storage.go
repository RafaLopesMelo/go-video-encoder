package vo

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

func NewStorage(input NewStorageDto) *Storage {
	storage := Storage{
		Provider:  input.Provider,
		Path:      input.Path,
		UploadURL: input.UploadURL,
	}

	return &storage
}
