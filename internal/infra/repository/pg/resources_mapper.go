package pg

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type persistenceResourceDto struct {
	id               string
	status           entity.ResourceStatus
	kind             entity.ResourceKind
	video_id         string
	storage_provider entity.ResourceStorageProvider
	size             int
	path             string
	upload_url       string
}

type resourcesMapper struct {
}

func (m *resourcesMapper) ToPersistence(resource entity.Resource) *persistenceResourceDto {
	return &persistenceResourceDto{
		id:               resource.ID.Value(),
		status:           resource.Status,
		kind:             resource.Kind,
		video_id:         resource.VideoID.Value(),
		storage_provider: resource.StorageProvider,
		size:             resource.Size,
		path:             resource.Path,
		upload_url:       resource.UploadURL,
	}
}

func (m *resourcesMapper) ToEntity(dto persistenceResourceDto) *entity.Resource {
	return &entity.Resource{
		ID:              vo.NewIDFromValue(dto.id),
		Status:          dto.status,
		Kind:            dto.kind,
		VideoID:         vo.NewIDFromValue(dto.video_id),
		StorageProvider: dto.storage_provider,
		Size:            dto.size,
		Path:            dto.path,
		UploadURL:       dto.upload_url,
	}
}

func newResourcesMapper() *resourcesMapper {
	return &resourcesMapper{}
}
