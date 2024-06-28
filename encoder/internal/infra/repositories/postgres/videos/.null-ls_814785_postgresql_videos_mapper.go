package mysql_videos_repository

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/videos"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects/unique_entity_id"
)

type PersistenceVideoDto struct {
	id          string
	resource_id string
	file_path   string
}

func (dto PersistenceVideoDto) ToEntity() *videos.Video {
	id := unique_entity_id.NewIDFromValue(dto.id)

	return videos.NewVideo(videos.NewVideoDto{
		FilePath:   dto.file_path,
		ResourceID: dto.resource_id,
	}, id)
}
