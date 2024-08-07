package entity_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
	"github.com/stretchr/testify/require"
)

func TestNewPendingJob(t *testing.T) {
	job := entity.NewJob(entity.NewJobDto{
		Kind:    entity.JobKindTranscode,
		VideoID: vo.NewID(),
	}, nil, nil)

	validated, err := entity.NewValidatedJob(*job)

	require.NoError(t, err)
	require.NotNil(t, validated)
	require.EqualValues(t, job.Status, entity.JobStatusPending)
}

func TestNewIdleJob(t *testing.T) {
	dependency := entity.NewJob(entity.NewJobDto{
		Kind:    entity.JobKindTranscode,
		VideoID: vo.NewID(),
	}, nil, nil)

	job := entity.NewJob(entity.NewJobDto{
		Kind:    entity.JobKindTranscode,
		VideoID: vo.NewID(),
	}, dependency, nil)

	validated, err := entity.NewValidatedJob(*job)

	require.NoError(t, err)
	require.NotNil(t, validated)
	require.EqualValues(t, job.Status, entity.JobStatusIdle)
}
