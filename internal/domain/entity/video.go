package entity

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type Video struct {
	ID         *vo.UniqueEntityID
	ResourceID string
	FilePath   string
}

type NewVideoDto struct {
	ResourceID string
	FilePath   string
}

func NewVideo(input NewVideoDto, id *vo.UniqueEntityID) *Video {
	if id == nil {
		id = vo.NewID()
	}

	video := Video{
		ID:         id,
		ResourceID: input.ResourceID,
		FilePath:   input.FilePath,
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
