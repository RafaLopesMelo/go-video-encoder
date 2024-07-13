package test

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects/unique_entity_id"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/common"
)

func DummyVideo() *entity.ValidatedVideo {
	id := unique_entity_id.NewID()

	video := entity.NewVideo(entity.NewVideoDto{
		ResourceID: "test",
		FilePath:   "/test",
	}, id)

	validated, err := entity.NewValidatedVideo(*video)

	if err != nil {
		panic("Dummy video not being built properly")
	}

	return validated
}

func DummyJob(videoId *unique_entity_id.UniqueEntityID) *entity.ValidatedJob {
	id := unique_entity_id.NewID()

	job := entity.NewJob(entity.NewJobDto{
		OutputBucketPath: "/",
		Status:           "PENDING",
		VideoID:          videoId,
		Error:            "",
	}, id)

	validated, err := entity.NewValidatedJob(*job)

	if err != nil {
		panic("Dummy job not being built properly")
	}

	return validated
}

type DummyController struct {
}

func (c *DummyController) Handle(request *common.Request) *common.Response {
	return &common.Response{}
}

func NewDummyController() *DummyController {
	return &DummyController{}
}
