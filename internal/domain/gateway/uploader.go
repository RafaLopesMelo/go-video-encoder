package gateway

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type PreparedUpload struct {
	URL      string
	Provider entity.ResourceStorageProvider
}

type Uploader interface {
	Prepare(videoID vo.UniqueEntityID) (PreparedUpload, error)
}
