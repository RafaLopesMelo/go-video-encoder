package command

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/repository"
)

type RegisterVideoUseCase struct {
	vr repository.VideosRepository
}

func (r RegisterVideoUseCase) Execute() error {
	video := entity.NewVideo(entity.NewVideoDto{}, nil)
	vv, err := entity.NewValidatedVideo(*video)

	if err != nil {
		return err
	}

	r.vr.Save(vv)

	return nil
}
