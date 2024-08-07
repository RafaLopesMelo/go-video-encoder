package pg

import (
	"database/sql"
	"fmt"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type persistenceJobDto struct {
	id            string
	status        entity.JobStatus
	kind          entity.JobKind
	video_id      string
	resource_id   sql.NullString
	depends_on_id sql.NullString
	error         string
}

type JobsRepo struct {
	connection *connection
}

func (r *JobsRepo) Save(validated entity.ValidatedJob) error {
	job := validated.Job()

	stmt := `
        INSERT INTO job
            (id, status, kind, video_id, resource_id, depends_on_id, error, created_at, updated_at)
        VALUES
            ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
        ON CONFLICT (id)
        DO UPDATE SET
            status = EXCLUDED.status,
            kind = EXCLUDED.kind,
            video_id = EXCLUDED.video_id,
            resource_id = EXCLUDED.resource_id,
            depends_on_id = EXCLUDED.depends_on_id,
            error = EXCLUDED.error,
            updated_at = NOW()
    `

	mapper := newJobMapper()
	data := mapper.ToPersistence(job)

	_, err := r.connection.DB.Exec(
		stmt,
		data.id,
		data.status,
		data.kind,
		data.video_id,
		data.resource_id,
		data.depends_on_id,
		data.error,
	)

	if err != nil {
		return fmt.Errorf("error saving job: %w", err)
	}

	return nil
}

func (r *JobsRepo) FindByID(id vo.UniqueEntityID) (entity.Job, error) {
	stmt := `
        SELECT id, status, kind, video_id, resource_id, depends_on_id, error  FROM job WHERE id = $1
    `

	result := r.connection.DB.QueryRow(stmt, id.Value())

	err := result.Err()

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Job{}, domainerrors.EntityNotFound
		}

		return entity.Job{}, fmt.Errorf("error finding job by id: %w", err)
	}

	dto := persistenceJobDto{}
	err = result.Scan(
		&dto.id,
		&dto.status,
		&dto.kind,
		&dto.video_id,
		&dto.resource_id,
		&dto.depends_on_id,
		&dto.error,
	)

	if err != nil {
		return entity.Job{}, fmt.Errorf("error scanning job: %w", err)
	}

	mapper := newJobMapper()
	entity := mapper.ToEntity(dto)

	return entity, nil
}

func NewJobsRepo(connection *connection) *JobsRepo {
	repository := JobsRepo{
		connection: connection,
	}

	return &repository
}
