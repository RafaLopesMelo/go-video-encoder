package postgres_jobs_repository

import (
	"database/sql"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repositories/postgres"
)

type PostgresJobsRepository struct {
	connection *postgres.Connection
}

func (repo *PostgresJobsRepository) Save(validated *entity.ValidatedJob) error {
	job := validated.Job()

	stmt := `
        INSERT INTO jobs
            (id, output_bucket_path, status, video_id, error, created_at, updated_at)
        VALUES
            ($1, $2, $3, $4, $5, NOW(), NOW())
        ON CONFLICT (id)
        DO UPDATE SET
            output_bucket_path = EXCLUDED.output_bucket_path,
            status = EXCLUDED.status,
            error = EXCLUDED.error,
            updated_at = NOW()
    `

	_, err := repo.connection.DB.Exec(
		stmt,
		job.ID.Value(),
		job.OutputBucketPath,
		job.Status,
		job.VideoID.Value(),
		job.Error,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *PostgresJobsRepository) FindByID(id vo.UniqueEntityID) (*entity.Job, error) {
	stmt := `
        SELECT id, output_bucket_path, status, video_id, error FROM jobs WHERE id = $1
    `

	result := repo.connection.DB.QueryRow(stmt, id.Value())

	err := result.Err()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewEntityNotFoundError()
		}

		return nil, err
	}

	dto := &PersistenceJobDto{}

	err = result.Scan(
		&dto.id,
		&dto.output_bucket_path,
		&dto.status,
		&dto.video_id,
		&dto.error,
	)

	if err != nil {
		return nil, err
	}

	return dto.ToEntity(), nil
}

func NewPostgresJobsRepository(connection *postgres.Connection) *PostgresJobsRepository {
	repository := PostgresJobsRepository{
		connection: connection,
	}

	return &repository
}
