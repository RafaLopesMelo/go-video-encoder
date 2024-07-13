package postgres_videos_repository

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects/unique_entity_id"
)

type PersistenceVideoDto struct {
	id          string
	resource_id string
	file_path   string
}

func (dto PersistenceVideoDto) ToEntity() *entity.Video {
	id := unique_entity_id.NewIDFromValue(dto.id)

	return entity.NewVideo(entity.NewVideoDto{
		FilePath:   dto.file_path,
		ResourceID: dto.resource_id,
	}, id)
}
