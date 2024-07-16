package pg

import (
	"database/sql"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	domainerrors "github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type VideosRepository struct {
	connection *connection
}

func (repo *VideosRepository) Save(validated *entity.ValidatedVideo) error {
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

	_, err := repo.connection.DB.Exec(
		stmt,
		video.ID.Value(),
		video.Status,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *VideosRepository) FindByID(id vo.UniqueEntityID) (*entity.Video, error) {
	stmt := `
        SELECT id, status FROM video WHERE id = $1
    `

	result := repo.connection.DB.QueryRow(stmt, id.Value())

	err := result.Err()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domainerrors.EntityNotFound
		}

		return nil, err
	}

	dto := &PersistenceVideoDto{}

	err = result.Scan(
		&dto.id,
		&dto.status,
	)

	if err != nil {
		return nil, err
	}

	video := dto.ToEntity()
	return video, nil
}

func NewVideosRepository(connection *connection) *VideosRepository {
	repository := VideosRepository{
		connection: connection,
	}

	return &repository
}
