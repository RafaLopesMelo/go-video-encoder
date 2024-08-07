package pg

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type persistenceVideoDto struct {
	id     string
	status entity.VideoStatus
}

type videoMapper struct {
}

func (m videoMapper) ToPersistence(entity entity.Video) persistenceVideoDto {
	dto := persistenceVideoDto{}

	dto.id = entity.ID.Value()
	dto.status = entity.Status

	return dto
}

func (m videoMapper) ToEntity(dto persistenceVideoDto) entity.Video {
	id := vo.NewIDFromValue(dto.id)

	return entity.LoadVideo(entity.LoadVideoDto{
		Status: dto.status,
	}, id)
}

func newVideoMapper() *videoMapper {
	return &videoMapper{}
}
