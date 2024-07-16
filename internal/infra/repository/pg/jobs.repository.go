package pg

import (
	"database/sql"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type JobsRepository struct {
	connection *connection
}

func (repo *JobsRepository) Save(validated *entity.ValidatedJob) error {
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

	dependsOnID := &sql.NullString{
		String: "",
		Valid:  false,
	}
	if job.DependsOnID != nil {
		dependsOnID.Valid = true
		dependsOnID.String = job.DependsOnID.Value()
	}

	resourceID := &sql.NullString{
		String: "",
		Valid:  false,
	}
	if job.DependsOnID != nil {
		resourceID.Valid = true
		resourceID.String = job.ResourceID.Value()
	}

	_, err := repo.connection.DB.Exec(
		stmt,
		job.ID.Value(),
		job.Status,
		job.Kind,
		job.VideoID.Value(),
		resourceID,
		dependsOnID,
		job.Error,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *JobsRepository) FindByID(id vo.UniqueEntityID) (*entity.Job, error) {
	stmt := `
        SELECT id, status, kind, video_id, resource_id, depends_on_id, error  FROM job WHERE id = $1
    `

	result := repo.connection.DB.QueryRow(stmt, id.Value())

	err := result.Err()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domainerrors.EntityNotFound
		}

		return nil, err
	}

	dto := &PersistenceJobDto{}

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
		return nil, err
	}

	return dto.ToEntity(), nil
}

func NewJobsRepository(connection *connection) *JobsRepository {
	repository := JobsRepository{
		connection: connection,
	}

	return &repository
}
