package entity

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects/unique_entity_id"
)

type Job struct {
	ID               *unique_entity_id.UniqueEntityID
	OutputBucketPath string
	Status           string
	VideoID          *unique_entity_id.UniqueEntityID
	Error            string
}

type NewJobDto struct {
	OutputBucketPath string
	Status           string
	VideoID          *unique_entity_id.UniqueEntityID
	Error            string
}

func NewJob(input NewJobDto, id *unique_entity_id.UniqueEntityID) *Job {
	if id == nil {
		id = unique_entity_id.NewID()
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
