package pg

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type ResourcesRepository struct {
	connection *connection
}

func (r *ResourcesRepository) Save(validated *entity.ValidatedResource) error {
	return nil
}

func (r *ResourcesRepository) FindByID(id vo.UniqueEntityID) (*entity.Resource, error) {
	return nil, nil
}

func NewResourcesRepository(connection *connection) *ResourcesRepository {
	repository := ResourcesRepository{
		connection: connection,
	}

	return &repository
}
