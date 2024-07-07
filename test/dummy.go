package test

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/jobs"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/videos"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects/unique_entity_id"
)

func DummyVideo() *videos.ValidatedVideo {
    id := unique_entity_id.NewID()

    video := videos.NewVideo(videos.NewVideoDto{
        ResourceID: "test",
        FilePath: "/test",
    }, id)

    validated, err := videos.NewValidatedVideo(*video)

    if err != nil {
        panic("Dummy video not being built properly")
    }

    return validated
}

func DummyJob(videoId *unique_entity_id.UniqueEntityID) *jobs.ValidatedJob {
    id := unique_entity_id.NewID()

    job := jobs.NewJob(jobs.NewJobDto{
        OutputBucketPath: "/",
        Status: "PENDING",
        VideoID: videoId,
        Error: "",
    }, id)

    validated, err := jobs.NewValidatedJob(*job)

    if err != nil {
        panic("Dummy job not being built properly")
    }

    return validated
}
