package gcp

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"cloud.google.com/go/storage"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/gateway"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/config/env"
	"github.com/google/uuid"
)

type Uploader struct{}

func (u Uploader) Prepare(videoID vo.UniqueEntityID) (gateway.PreparedUpload, error) {
	bucket := env.Get("GCP_BUCKET")

	uuid, _ := uuid.NewV7()
	obj := videoID.Value() + "/" + uuid.String()

	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		return gateway.PreparedUpload{}, fmt.Errorf("error creating GCP storage client: %w", err)
	}

	defer client.Close()

	params := url.Values{}
	params.Add("uploadType", "resumable")

	opts := &storage.SignedURLOptions{
		Scheme:          storage.SigningSchemeV4,
		Method:          "POST",
		Expires:         time.Now().Add(time.Hour * 24),
		QueryParameters: params,
	}

	url, err := client.Bucket(bucket).SignedURL(obj, opts)

	if err != nil {
		return gateway.PreparedUpload{}, fmt.Errorf("error creating signed URL: %w", err)
	}

	return gateway.PreparedUpload{
		URL:      url,
		Provider: entity.ResourceStorageProviderGCP,
		Path:     obj,
	}, nil
}

func NewUploader() Uploader {
	return Uploader{}
}
