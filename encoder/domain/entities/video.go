package entities

import (
	"time"

	"github.com/RafaLopesMelo/go-video-encoder/domain/value_objects"
)

type Video struct {
    ID          value_objects.UniqueEntityID
    ResourceID  string
    FilePath    string
    CreatedAt   time.Time
}

func NewVideo() *Video {
    video := Video{}
    return &video
}

func (video *Video) Validate() error {
    return nil
}
