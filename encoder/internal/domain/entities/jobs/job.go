package jobs

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects"
)

type Job struct {
    ID                  *value_objects.UniqueEntityID
    OutputBucketPath    string
    Status              string
    VideoID             *value_objects.UniqueEntityID
    Error               string
}

type InputJobDto struct {
    OutputBucketPath    string
    Status              string
    VideoID             *value_objects.UniqueEntityID
    Error               string
}

func NewJob(input InputJobDto, id *value_objects.UniqueEntityID) *Job {
    if id == nil {
        id = value_objects.NewID()
    }

    job := Job{
        ID: id,
        OutputBucketPath: input.OutputBucketPath,
        Status: input.Status,
        VideoID: input.VideoID,
        Error: input.Error,
    }

    return &job
}

func (job *Job) validate() error {
    if job.VideoID == nil  {
        return errors.NewRequiredPropertyError("VideoID")
    }

    if job.OutputBucketPath == "" {
        return errors.NewRequiredPropertyError("OutputBucketPath")
    }

    return nil
}
