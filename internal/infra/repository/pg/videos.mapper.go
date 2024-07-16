package pg

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type PersistenceVideoDto struct {
	id     string
	status entity.VideoStatus
}

func (dto PersistenceVideoDto) ToEntity() *entity.Video {
	id := vo.NewIDFromValue(dto.id)

	return entity.LoadVideo(entity.LoadVideoDto{
		Status: dto.status,
	}, id)
}
