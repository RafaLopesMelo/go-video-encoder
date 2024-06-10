package videos

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects"
)

type Video struct {
    ID          *value_objects.UniqueEntityID
    ResourceID  string
    FilePath    string
}

type NewVideoDto struct {
    ResourceID string
    FilePath string
}

func NewVideo(input NewVideoDto, id *value_objects.UniqueEntityID) *Video {
    if id == nil {
        id = value_objects.NewID()
    }

    video := Video{
        ID: id,
        ResourceID: input.ResourceID,
        FilePath: input.FilePath,
    }

    return &video
}

func (video *Video) validate() error {
    if video.FilePath == "" {
        return errors.NewRequiredPropertyError("FilePath")
    }

    if video.ResourceID == "" {
        return errors.NewRequiredPropertyError("ResourceID")
    }

    return nil
}
