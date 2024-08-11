package entity

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type JobStatus string

const (
	JobStatusIdle    JobStatus = "IDLE"
	JobStatusPending JobStatus = "PENDING"
	JobStatusRunning JobStatus = "RUNNING"
	JobStatusDone    JobStatus = "DONE"
	JobStatusFailed  JobStatus = "FAILED"
)

type JobKind string

const (
	JobKindTranscode JobKind = "TRANSCODE"
)

type Job struct {
	ID          *vo.UniqueEntityID
	Status      JobStatus
	Kind        JobKind
	VideoID     *vo.UniqueEntityID
	ResourceID  *vo.UniqueEntityID
	DependsOnID *vo.UniqueEntityID
	Error       string
}

type NewJobDto struct {
	Kind    JobKind
	VideoID *vo.UniqueEntityID
}

func NewJob(input NewJobDto, dependsOn *Job, id *vo.UniqueEntityID) *Job {
	if id == nil {
		id = vo.NewID()
	}

	var status JobStatus = JobStatusPending
	if dependsOn != nil && dependsOn.Status != JobStatusDone {
		status = JobStatusIdle
	}

	var depensOnID *vo.UniqueEntityID = nil
	if dependsOn != nil {
		depensOnID = dependsOn.ID
	}

	job := Job{
		ID:          id,
		Status:      status,
		Kind:        input.Kind,
		VideoID:     input.VideoID,
		ResourceID:  nil,
		DependsOnID: depensOnID,
		Error:       "",
	}

	return &job
}

type LoadJobDto struct {
	Status      JobStatus
	Kind        JobKind
	VideoID     *vo.UniqueEntityID
	ResourceID  *vo.UniqueEntityID
	DependsOnID *vo.UniqueEntityID
	Error       string
}

func LoadJob(input LoadJobDto, id *vo.UniqueEntityID) Job {
	job := Job{
		ID:          id,
		Status:      input.Status,
		Kind:        input.Kind,
		VideoID:     input.VideoID,
		ResourceID:  input.ResourceID,
		DependsOnID: input.DependsOnID,
		Error:       input.Error,
	}

	return job
}

func (job *Job) validate() error {
	if job.VideoID == nil {
		return domainerrors.RequiredProperty
	}

	return nil
}
