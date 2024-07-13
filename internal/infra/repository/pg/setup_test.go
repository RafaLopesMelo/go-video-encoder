package pg_test

import (
	"os"
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/config/env"
)

func TestMain(m *testing.M) {
	err := env.Load(".env")

	if err != nil {
		panic(err)
	}

	jobTestSetup()

	os.Exit(m.Run())
}
