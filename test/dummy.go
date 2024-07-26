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
		Kind:    entity.JobKindTranscode,
		VideoID: videoId,
	}, nil, id)

	validated, err := entity.NewValidatedJob(*job)

	if err != nil {
		panic("Dummy job not being built properly")
	}

	return validated
}

func DummyResource(videoId *vo.UniqueEntityID) *entity.ValidatedResource {
	id := vo.NewID()

	resource := entity.NewResource(entity.NewResourceDto{
		Kind:            entity.ResourceKindRawVideo,
		VideoID:         videoId,
		StorageProvider: entity.ResourceStorageProviderGCP,
		Path:            "/test",
		UploadURL:       "/test",
		Size:            100,
	}, id)

	validated, err := entity.NewValidatedResource(*resource)

	if err != nil {
		panic("Dummy resource not being built properly")
	}

	return validated
}
