package pg

import (
	"database/sql"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type VideosRepository struct {
	connection *connection
}

func (repo *VideosRepository) Save(validated *entity.ValidatedVideo) error {
	video := validated.Video()

	stmt := `
        INSERT INTO videos
            (id, resource_id, file_path, created_at, updated_at)
        VALUES
            ($1, $2, $3, NOW(), NOW())
        ON CONFLICT (id)
        DO UPDATE SET
            resource_id = EXCLUDED.resource_id,
            file_path = EXCLUDED.file_path,
            updated_at = NOW()
        ;
    `

	_, err := repo.connection.DB.Exec(
		stmt,
		video.ID.Value(),
		video.ResourceID,
		video.FilePath,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *VideosRepository) FindByID(id vo.UniqueEntityID) (*entity.Video, error) {
	stmt := `
        SELECT id, resource_id, file_path FROM videos WHERE id = $1
    `

	result := repo.connection.DB.QueryRow(stmt, id.Value())

	err := result.Err()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewEntityNotFoundError()
		}

		return nil, err
	}

	dto := &PersistenceVideoDto{}

	err = result.Scan(
		&dto.id,
		&dto.resource_id,
		&dto.file_path,
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
