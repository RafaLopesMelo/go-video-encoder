package repositories

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/jobs"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects/unique_entity_id"
)

type JobsRepository interface {
	Save(job *jobs.ValidatedJob) error
	findByID(id value_objects.UniqueEntityID) (*jobs.Job, error)
}
