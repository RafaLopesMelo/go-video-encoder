package repositories

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects/unique_entity_id"
)

type JobsRepository interface {
	Save(job *entity.ValidatedJob) error
	findByID(id unique_entity_id.UniqueEntityID) (*entity.Job, error)
}
