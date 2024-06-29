package env_test

import (
	"os"
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/configs/env"
	"github.com/stretchr/testify/require"
)

func TestLoadEnv(t *testing.T) {
    file, err := os.CreateTemp(".", ".env-*")

    if err != nil {
        t.Errorf("Could not create env temp file")
        return
    }

    defer os.Remove(file.Name())

    key := "TEST_VAR"
    value := "testing"

    file.WriteString("TEST_VAR=testing")
    err = env.Load(file.Name())

    require.Nil(t, err)
    require.Equal(t, env.Get(key), value)
}

func TestLoadEnvWhenFileDoNotExists(t *testing.T) {
    err := env.Load("invalid_file")
    require.NotNil(t, err)
}
