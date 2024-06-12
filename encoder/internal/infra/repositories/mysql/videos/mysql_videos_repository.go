package mysql_videos_repository

import (
	"database/sql"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/videos"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects/unique_entity_id"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repositories/mysql"
)

type MySqlVideosRepository struct {
	connection *mysql.Connection
}

func (repo *MySqlVideosRepository) Save(validated *videos.ValidatedVideo) error {
	video := validated.Video()

	stmt := `
        INSERT INTO videos
            (id, resource_id, file_path, created_at, updated_at)
        VALUES
            (?, ?, ?, NOW(), NOW())
        ON DUPLICATE KEY UPDATE
            resource_id = ?,
            file_path = ?,
            updated_at = NOW()
        ;
    `

	_, err := repo.connection.DB.Exec(
		stmt,
		video.ID.Value(),
		video.ResourceID,
		video.FilePath,
		video.ResourceID,
		video.FilePath,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *MySqlVideosRepository) FindByID(id unique_entity_id.UniqueEntityID) (*videos.Video, error) {
	stmt := `
        SELECT id, resource_id, file_path FROM videos WHERE id = ?
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

func NewMysqlVideosRepository(connection *mysql.Connection) *MySqlVideosRepository {
	repository := MySqlVideosRepository{
		connection: connection,
	}

	return &repository
}
