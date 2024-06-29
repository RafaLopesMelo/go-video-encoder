package postgres_jobs_repository

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/jobs"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repositories/postgres"
)

type PostgresJobsRepository struct {
	connection *postgres.Connection
}

func (repo *PostgresJobsRepository) Save(job *jobs.ValidatedJob) () {

}
