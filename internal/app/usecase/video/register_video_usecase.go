package video

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/gateway"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/repository"
)

type RegisterUseCase struct {
	vr       repository.VideosRepository
	uploader gateway.Uploader
}

func (r RegisterUseCase) Execute() error {
	video := entity.NewVideo(entity.NewVideoDto{}, nil)
	vv, err := entity.NewValidatedVideo(*video)

	if err != nil {
		return err
	}

	prepared, err := r.uploader.Prepare(*vv.Video().ID)

	if err != nil {
		return err
	}

	raw := entity.NewRawVideo(entity.NewRawVideoDto{
		NewResourceDto: entity.NewResourceDto{
			VideoID:         vv.Video().ID,
			UploadURL:       prepared.URL,
			StorageProvider: prepared.Provider,
		},
	}, nil)

	r.vr.Save(vv)

	return nil
}
