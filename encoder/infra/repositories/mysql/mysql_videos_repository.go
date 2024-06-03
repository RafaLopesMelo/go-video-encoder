package mysql

import (
	"github.com/RafaLopesMelo/go-video-encoder/domain/entities"
)

type MySqlVideosRepository struct {
    connection *Connection
}

func (repo *MySqlVideosRepository) Save() error {
    return nil
}

func (repo *MySqlVideosRepository) FindByID() *entities.Video {
    return entities.NewVideo()
}

func NewMysqlVideosRepository(connection *Connection) *MySqlVideosRepository {
    repository := MySqlVideosRepository{
        connection: connection,
    }

    return &repository
}
