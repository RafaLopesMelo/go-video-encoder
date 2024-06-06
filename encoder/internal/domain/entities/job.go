package entities

import (
	"time"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects"
)

type Job struct {
    ID                  value_objects.UniqueEntityID
    OutputBucketPath    string
    Status              string
    Video               *Video
    Error               string
    CreatedAt           time.Time
    UpdatedAt           time.Time   
}

func NewJob(output string, status string, video *Video) (*Job, error) {
    job := Job{
        OutputBucketPath: output,
        Status: status,
        Video: video,
    }

    job.prepare()

    err := job.Validate()

    if err != nil {
        return nil, err
    }

    return &job, nil
}

func (job *Job) prepare() {
    job.ID = *value_objects.NewID()
    job.CreatedAt = time.Now()
    job.UpdatedAt = time.Now()
}

func (job *Job) Validate() error {
    return nil
}
