package test

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

func DummyVideo() *entity.ValidatedVideo {
	id := vo.NewID()

	video := entity.NewVideo(entity.NewVideoDto{}, id)

	validated, err := entity.NewValidatedVideo(*video)

	if err != nil {
		panic("Dummy video not being built properly")
	}

	return validated
}

func DummyJob(videoId *vo.UniqueEntityID) *entity.ValidatedJob {
	id := vo.NewID()

	job := entity.NewJob(entity.NewJobDto{
		Type:    entity.JobTypeTranscode,
		VideoID: vo.NewID(),
	}, nil, id)

	validated, err := entity.NewValidatedJob(*job)

	if err != nil {
		panic("Dummy job not being built properly")
	}

	return validated
}

func DummyStorage() *vo.Storage {
	storage, err := vo.NewStorage(vo.NewStorageDto{
		Provider:  vo.StorageProviderGCP,
		Path:      "/test",
		UploadURL: "test",
	})

	if err != nil {
		panic("Dummy storage not being built properly")
	}

	return storage
}
