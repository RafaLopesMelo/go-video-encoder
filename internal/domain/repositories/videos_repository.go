package repositories

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/videos"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects/unique_entity_id"
)

type VideosRepository interface {
    Save(video *videos.ValidatedVideo) error
	findByID(id unique_entity_id.UniqueEntityID) (*videos.Video, error)
}
