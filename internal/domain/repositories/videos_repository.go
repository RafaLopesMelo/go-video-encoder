package repositories

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects/unique_entity_id"
)

type VideosRepository interface {
	Save(video *entity.ValidatedVideo) error
	findByID(id unique_entity_id.UniqueEntityID) (*entity.Video, error)
}
