package entity

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type Job struct {
	ID               *vo.UniqueEntityID
	OutputBucketPath string
	Status           string
	VideoID          *vo.UniqueEntityID
	Error            string
}

type NewJobDto struct {
	OutputBucketPath string
	Status           string
	VideoID          *vo.UniqueEntityID
	Error            string
}

func NewJob(input NewJobDto, id *vo.UniqueEntityID) *Job {
	if id == nil {
		id = vo.NewID()
	}

	job := Job{
		ID:               id,
		OutputBucketPath: input.OutputBucketPath,
		Status:           input.Status,
		VideoID:          input.VideoID,
		Error:            input.Error,
	}

	return &job
}

func (job *Job) validate() error {
	if job.VideoID == nil {
		return errors.NewRequiredPropertyError("VideoID")
	}

	if job.OutputBucketPath == "" {
		return errors.NewRequiredPropertyError("OutputBucketPath")
	}

	return nil
}
