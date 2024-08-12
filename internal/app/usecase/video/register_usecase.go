package video

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/gateway"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/repo"
)

type RegisterUseCase struct {
	vr repo.VideosRepository
	rr repo.ResourcesRepository
	u  gateway.Uploader
}

type CreatedVideo struct {
	ID        string
	UploadURL string
}

func (r RegisterUseCase) Execute() (CreatedVideo, error) {
	video := entity.NewVideo(entity.NewVideoDto{}, nil)
	vv, err := entity.NewValidatedVideo(video)

	if err != nil {
		return CreatedVideo{}, err
	}

	raw := entity.NewRawVideo(entity.NewRawVideoDto{
		NewResourceDto: entity.NewResourceDto{
			VideoID: vv.Video().ID,
		},
	}, nil)
	vr, err := entity.NewValidatedResource(raw)

	if err != nil {
		return CreatedVideo{}, err
	}

	prepared, err := r.u.Prepare(*video.ID)

	if err != nil {
		return CreatedVideo{}, err
	}

	r.vr.Save(vv)
	r.rr.Save(vr)

	return CreatedVideo{
		ID:        vv.Video().ID.Value(),
		UploadURL: prepared.URL,
	}, nil
}

func NewRegisterUseCase(vr repo.VideosRepository, rr repo.ResourcesRepository, u gateway.Uploader) RegisterUseCase {
	return RegisterUseCase{
		vr: vr,
		rr: rr,
		u:  u,
	}
}
