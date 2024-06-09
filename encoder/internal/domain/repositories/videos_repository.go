package repositories

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/videos"
)

type VideosRepository interface {
    Save(video *videos.Video) error
	findByID() *videos.Video
}
