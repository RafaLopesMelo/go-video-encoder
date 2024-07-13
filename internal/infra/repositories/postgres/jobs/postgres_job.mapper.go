package postgres_jobs_repository

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects/unique_entity_id"
)

type PersistenceJobDto struct {
	id                 string
	output_bucket_path string
	status             string
	video_id           string
	error              string
}

func (dto PersistenceJobDto) ToEntity() *entity.Job {
	id := unique_entity_id.NewIDFromValue(dto.id)
	videoId := unique_entity_id.NewIDFromValue(dto.video_id)

	return entity.NewJob(entity.NewJobDto{
		Status:           dto.status,
		Error:            dto.error,
		VideoID:          videoId,
		OutputBucketPath: dto.output_bucket_path,
	}, id)
}
