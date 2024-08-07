package gcp_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/config/env"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/gateway/gcp"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	err := env.Load(".env")

	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestPrepare(t *testing.T) {
	fmt.Println("TestPrepare")
	u := gcp.NewUploader()

	videoID := vo.NewID()

	prepared, err := u.Prepare(*videoID)
	require.Nil(t, err)
	require.NotEmpty(t, prepared.URL)
	require.Equal(t, prepared.Provider, entity.ResourceStorageProviderGCP)
}
