package test

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entities/videos"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects"
)

func DummyVideo() *videos.ValidatedVideo {
    id := value_objects.NewID()

    video := videos.NewVideo(videos.NewVideoDto{
        ResourceID: "test",
        FilePath: "/test",
    }, id)

    validated, err := videos.NewValidatedVideo(*video)

    if err != nil {
        panic("Dummy video not being build properly")
    }

    return validated
}
