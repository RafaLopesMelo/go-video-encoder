package entity_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	job := entity.NewJob(entity.NewJobDto{
		OutputBucketPath: "/test",
		VideoID:          vo.NewID(),
		Status:           "PENDING",
		Error:            "",
	}, nil)

	validated, err := entity.NewValidatedJob(*job)

	require.Nil(t, err)
	require.NotNil(t, validated)
}

func TestNewJobWithoutOutputBucketPath(t *testing.T) {
	job := entity.NewJob(entity.NewJobDto{
		OutputBucketPath: "",
		VideoID:          vo.NewID(),
		Status:           "PENDING",
		Error:            "",
	}, nil)

	_, err := entity.NewValidatedJob(*job)

	expected := errors.NewRequiredPropertyError("OutputBucketPath")

	require.Error(t, err, expected.Error())
}

func TestNewJobWithoutVideoID(t *testing.T) {
	job := entity.NewJob(entity.NewJobDto{
		OutputBucketPath: "/test",
		VideoID:          nil,
		Status:           "PENDING",
		Error:            "",
	}, nil)

	_, err := entity.NewValidatedJob(*job)

	expected := errors.NewRequiredPropertyError("VideoID")

	require.Error(t, err, expected.Error())
}
