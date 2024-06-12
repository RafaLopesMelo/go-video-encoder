package repositories

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/jobs"
)

type JobsRepository interface {
	Save(job *jobs.ValidatedJob) error
	findByID() *jobs.Job
}
