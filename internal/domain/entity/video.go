package entity

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type VideoStatus string

const (
	VideoStatusPending    VideoStatus = "PENDING"
	VideoStatusUploaded   VideoStatus = "UPLOADED"
	VideoStatusProcessing VideoStatus = "PROCESSING"
	VideoStatusProcessed  VideoStatus = "PROCESSED"
	VideoStatusFailed     VideoStatus = "FAILED"
)

type Video struct {
	ID     *vo.UniqueEntityID
	Status VideoStatus
}

type NewVideoDto struct {
}

func NewVideo(input NewVideoDto, id *vo.UniqueEntityID) *Video {
	if id == nil {
		id = vo.NewID()
	}

	video := Video{
		ID:     id,
		Status: VideoStatusPending,
	}

	return &video
}

type LoadVideoDto struct {
	Status VideoStatus
}

func LoadVideo(input LoadVideoDto, id *vo.UniqueEntityID) Video {
	video := Video{
		ID:     id,
		Status: input.Status,
	}

	return video
}

func (video *Video) validate() error {
	return nil
}
