package pg

import (
	"fmt"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type ResourcesRepository struct {
	connection *connection
}

func (r *ResourcesRepository) Save(validated *entity.ValidatedResource) error {
	rw := validated.Wrapper()

	stmt := `
        INSERT INTO resource
            (id, status, kind, video_id, storage_provider, size, path, upload_url, metadata, created_at, updated_at)
        VALUES
            ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())
        ON CONFLICT (id)
        DO UPDATE SET
            status = EXCLUDED.status,
            kind = EXCLUDED.kind,
            video_id = EXCLUDED.video_id,
            storage_provider = EXCLUDED.storage_provider,
            size = EXCLUDED.size,
            path = EXCLUDED.path,
            upload_url = EXCLUDED.upload_url,
            metadata = EXCLUDED.metadata,
            updated_at = NOW()
    `

	mapper := newResourcesMapper()
	data := mapper.ToPersistence(rw)

	_, err := r.connection.DB.Exec(
		stmt,
		data.id,
		data.status,
		data.kind,
		data.video_id,
		data.storage_provider,
		data.size,
		data.path,
		data.upload_url,
		data.metadata,
	)

	if err != nil {
		return fmt.Errorf("error saving resource: %w", err)
	}

	return nil
}

func (r *ResourcesRepository) FindByID(id vo.UniqueEntityID) (entity.ResourceWrapper, error) {
	stmt := `
        SELECT id, status, kind, video_id, storage_provider, size, path, upload_url, metadata FROM resource WHERE id = $1
    `

	result := r.connection.DB.QueryRow(stmt, id.Value())

	dto := persistenceResourceDto{}

	err := result.Scan(
		&dto.id,
		&dto.status,
		&dto.kind,
		&dto.video_id,
		&dto.storage_provider,
		&dto.size,
		&dto.path,
		&dto.upload_url,
		&dto.metadata,
	)

	if err != nil {
		return nil, fmt.Errorf("error scanning resource: %w", err)
	}

	mapper := newResourcesMapper()
	entity, err := mapper.ToEntity(dto)

	if err != nil {
		return nil, fmt.Errorf("error mapping resource: %w", err)
	}

	return entity, nil
}

func NewResourcesRepository(connection *connection) *ResourcesRepository {
	repository := ResourcesRepository{
		connection: connection,
	}

	return &repository
}
