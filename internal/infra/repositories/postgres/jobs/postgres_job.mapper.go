package postgres_jobs_repository

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/jobs"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects/unique_entity_id"
)

type PersistenceJobDto struct {
	id                 string
	output_bucket_path string
	status             string
	video_id           string
	error              string
}

func (dto PersistenceJobDto) ToEntity() *jobs.Job {
	id := unique_entity_id.NewIDFromValue(dto.id)
	videoId := unique_entity_id.NewIDFromValue(dto.video_id)

	return jobs.NewJob(jobs.NewJobDto{
		Status:           dto.status,
		Error:            dto.error,
		VideoID:          videoId,
		OutputBucketPath: dto.output_bucket_path,
	}, id)
}
