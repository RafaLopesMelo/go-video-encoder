package pg

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	domainerrors "github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
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
	metadata         *jsonable
}

type resourcesMapper struct {
}

func (m *resourcesMapper) ToPersistence(rw entity.ResourceWrapper) persistenceResourceDto {
	r := rw.Resource()

	jsonable := NewJSONable(rw.Metadata())

	return persistenceResourceDto{
		id:               r.ID.Value(),
		status:           r.Status,
		kind:             r.Kind,
		video_id:         r.VideoID.Value(),
		storage_provider: r.StorageProvider,
		size:             r.Size,
		path:             r.Path,
		upload_url:       r.UploadURL,
		metadata:         jsonable,
	}
}

func (m *resourcesMapper) ToEntity(dto persistenceResourceDto) (entity.ResourceWrapper, error) {
	id := vo.NewIDFromValue(dto.id)
	resource := entity.NewResourceDto{
		VideoID:         vo.NewIDFromValue(dto.video_id),
		StorageProvider: dto.storage_provider,
		Size:            dto.size,
		Path:            dto.path,
		UploadURL:       dto.upload_url,
	}

	metadata := dto.metadata.ToMap()

	if dto.kind == entity.ResourceKindRawVideo {
		rv := entity.NewRawVideo(entity.NewRawVideoDto{
			NewResourceDto: resource,
			Extension:      metadata["extension"].(string),
		}, id)

		rv.Resource().Status = dto.status
		return rv, nil
	}

	return nil, domainerrors.InvalidResourceKind
}

func newResourcesMapper() *resourcesMapper {
	return &resourcesMapper{}
}
