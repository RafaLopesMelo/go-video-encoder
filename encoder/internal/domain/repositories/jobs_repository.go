package repositories

import "github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities"

type JobsRepository interface {
	Save(job *entities.Job) error
	findByID() *entities.Job
}
