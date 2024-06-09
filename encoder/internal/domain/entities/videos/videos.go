package videos

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/value_objects"
)

type Video struct {
    ID          value_objects.UniqueEntityID
    ResourceID  string
    FilePath    string
}

type NewVideoDto struct {
    ResourceID string
    FilePath string
}

func NewVideo(input NewVideoDto, id *value_objects.UniqueEntityID) (*Video, error) {
    if id == nil {
        id = value_objects.NewID()
    }

    video := Video{
        ID: *id,
        ResourceID: input.ResourceID,
        FilePath: input.FilePath,
    }

    err := video.validate()

    if err != nil {
        return nil, err
    }

    return &video, nil
}

func (video *Video) validate() error {
    return nil
}
