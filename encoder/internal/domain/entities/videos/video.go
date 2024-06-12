package videos

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects/unique_entity_id"
)

type Video struct {
    ID          *unique_entity_id.UniqueEntityID
    ResourceID  string
    FilePath    string
}

type NewVideoDto struct {
    ResourceID string
    FilePath string
}

func NewVideo(input NewVideoDto, id *unique_entity_id.UniqueEntityID) *Video {
    if id == nil {
        id = unique_entity_id.NewID()
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
