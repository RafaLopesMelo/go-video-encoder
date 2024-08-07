package repo

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type ResourcesRepository interface {
	Save(resource entity.ValidatedResource) error
	FindByID(id vo.UniqueEntityID) (entity.ResourceWrapper, error)
}
