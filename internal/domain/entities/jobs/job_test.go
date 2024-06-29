package jobs_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/jobs"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects/unique_entity_id"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
    job := jobs.NewJob(jobs.InputJobDto{
        OutputBucketPath: "/test",
        VideoID: unique_entity_id.NewID(),
        Status: "PENDING",
        Error: "",
    }, nil)

    validated, err := jobs.NewValidatedJob(*job)

    require.Nil(t, err)
    require.NotNil(t, validated)
}

func TestNewJobWithoutOutputBucketPath(t *testing.T) {
    job := jobs.NewJob(jobs.InputJobDto{
        OutputBucketPath: "",
        VideoID: unique_entity_id.NewID(),
        Status: "PENDING",
        Error: "",
    }, nil)

    _, err := jobs.NewValidatedJob(*job)

    expected := errors.NewRequiredPropertyError("OutputBucketPath")

    require.Error(t, err, expected.Error())
}

func TestNewJobWithoutVideoID(t *testing.T) {
    job := jobs.NewJob(jobs.InputJobDto{
        OutputBucketPath: "/test",
        VideoID: nil,
        Status: "PENDING",
        Error: "",
    }, nil)

    _, err := jobs.NewValidatedJob(*job)

    expected := errors.NewRequiredPropertyError("VideoID")

    require.Error(t, err, expected.Error())
}
