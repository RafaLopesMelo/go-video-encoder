package repositories

import "github.com/RafaLopesMelo/go-video-encoder/domain/entities"

type VideosRepository interface {
    Save(video *entities.Video) error
	findByID() *entities.Video
}
