package repositories

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type VideosRepository interface {
	Save(video *entity.ValidatedVideo) error
	findByID(id vo.UniqueEntityID) (*entity.Video, error)
}
