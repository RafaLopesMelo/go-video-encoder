package mysql

import (
	"github.com/RafaLopesMelo/go-video-encoder/domain/entities"
)

type MySqlVideosRepository struct {
    connection *Connection
}

func (r *MySqlVideosRepository) findByID() *entities.Video {
    return entities.NewVideo()
}

func NewMysqlVideosRepository(connection *Connection) *MySqlVideosRepository {
    repository := MySqlVideosRepository{
        connection: connection,
    }

    return &repository
}
