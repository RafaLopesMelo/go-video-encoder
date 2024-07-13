package postgres_jobs_repository_test

import (
	"os"
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/configs/env"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repositories/postgres"
	postgres_jobs_repository "github.com/RafaLopesMelo/go-video-encoder/internal/infra/repositories/postgres/jobs"
	postgres_videos_repository "github.com/RafaLopesMelo/go-video-encoder/internal/infra/repositories/postgres/videos"
	"github.com/RafaLopesMelo/go-video-encoder/test"
	"github.com/stretchr/testify/require"
)

var videoId *vo.UniqueEntityID

func TestMain(m *testing.M) {
	err := env.Load(".env")

	if err != nil {
		panic(err)
	}

	connection := postgres.NewConnection()
	repo := postgres_videos_repository.NewPostgresVideosRepository(connection)

	video := test.DummyVideo()
	repo.Save(video)
	videoId = video.Video().ID

	os.Exit(m.Run())
}

func TestCreateJob(t *testing.T) {
	connection := postgres.NewConnection()
	repo := postgres_jobs_repository.NewPostgresJobsRepository(connection)

	dummy := test.DummyJob(videoId)

	repo.Save(dummy)
	job, err := repo.FindByID(*dummy.Job().ID)

	require.Nil(t, err)
	require.Equal(t, dummy.Job().ID, job.ID)
}

func TestUpdateVideo(t *testing.T) {
	connection := postgres.NewConnection()
	repo := postgres_jobs_repository.NewPostgresJobsRepository(connection)

	dummy := test.DummyJob(videoId)
	repo.Save(dummy)

	job := dummy.Job()
	job.OutputBucketPath = "/test-updated"
	validated, _ := entity.NewValidatedJob(job)

	repo.Save(validated)
	updated, err := repo.FindByID(*job.ID)

	require.Nil(t, err)
	require.Equal(t, job.ID, updated.ID)
	require.Equal(t, job.OutputBucketPath, updated.OutputBucketPath)
}
