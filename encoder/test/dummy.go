package test

import (
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
