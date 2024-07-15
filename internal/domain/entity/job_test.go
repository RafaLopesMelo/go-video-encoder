package entity_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
	"github.com/stretchr/testify/require"
)

func TestNewPendingJob(t *testing.T) {
	job := entity.NewJob(entity.NewJobDto{
		Type:    entity.JobTypeTranscode,
		VideoID: vo.NewID(),
	}, nil, nil)

	validated, err := entity.NewValidatedJob(*job)

	require.Nil(t, err)
	require.NotNil(t, validated)
	require.EqualValues(t, entity.JobStatusPending, job.Status)
}

func TestNewIdleJob(t *testing.T) {
	dependency := entity.NewJob(entity.NewJobDto{
		Type:    entity.JobTypeTranscode,
		VideoID: vo.NewID(),
	}, nil, nil)

	job := entity.NewJob(entity.NewJobDto{
		Type:    entity.JobTypeTranscode,
		VideoID: vo.NewID(),
	}, dependency, nil)

	validated, err := entity.NewValidatedJob(*job)

	require.Nil(t, err)
	require.NotNil(t, validated)
	require.EqualValues(t, entity.JobStatusIdle, job.Status)
}
