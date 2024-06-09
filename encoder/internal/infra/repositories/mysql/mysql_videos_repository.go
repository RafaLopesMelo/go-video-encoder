package mysql

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/videos"
)

type MySqlVideosRepository struct {
	connection *Connection
}

func (repo *MySqlVideosRepository) Save() error {
	return nil
}

    func (repo *MySqlVideosRepository) FindByID() *videos.Video {
    return nil
}

func NewMysqlVideosRepository(connection *Connection) *MySqlVideosRepository {
	repository := MySqlVideosRepository{
		connection: connection,
	}

	return &repository
}
