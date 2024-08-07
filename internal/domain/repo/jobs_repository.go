package repo

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type JobsRepository interface {
	Save(job entity.ValidatedJob) error
	FindByID(id vo.UniqueEntityID) (entity.Job, error)
}
