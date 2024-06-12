package mysql_jobs_repository

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/jobs"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repositories/mysql"
)

type MySqlJobsRepository struct {
	connection *mysql.Connection
}

func (repo *MySqlJobsRepository) Save(job *jobs.ValidatedJob) () {

}
