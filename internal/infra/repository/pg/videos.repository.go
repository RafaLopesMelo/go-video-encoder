package pg

import (
	"database/sql"
	"fmt"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	domainerrors "github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type VideosRepository struct {
	connection *connection
}

func (r *VideosRepository) Save(validated *entity.ValidatedVideo) error {
	video := validated.Video()

	stmt := `
        INSERT INTO video
            (id, status, created_at, updated_at)
        VALUES
            ($1, $2, NOW(), NOW())
        ON CONFLICT (id)
        DO UPDATE SET
            status = EXCLUDED.status,
            updated_at = NOW()
        ;
    `

	mapper := newVideoMapper()
	data := mapper.ToPersistence(video)

	_, err := r.connection.DB.Exec(
		stmt,
		data.id,
		data.status,
	)

	if err != nil {
		return fmt.Errorf("error saving video: %w", err)
	}

	return nil
}

func (r *VideosRepository) FindByID(id vo.UniqueEntityID) (*entity.Video, error) {
	stmt := `
        SELECT id, status FROM video WHERE id = $1
    `

	result := r.connection.DB.QueryRow(stmt, id.Value())

	err := result.Err()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domainerrors.EntityNotFound
		}

		return nil, fmt.Errorf("error finding video by id: %w", err)
	}

	dto := persistenceVideoDto{}
	err = result.Scan(
		&dto.id,
		&dto.status,
	)

	if err != nil {
		return nil, fmt.Errorf("error scanning video: %w", err)
	}

	mapper := newVideoMapper()
	entity := mapper.ToEntity(dto)

	return entity, nil
}

func NewVideosRepository(connection *connection) *VideosRepository {
	repository := VideosRepository{
		connection: connection,
	}

	return &repository
}
