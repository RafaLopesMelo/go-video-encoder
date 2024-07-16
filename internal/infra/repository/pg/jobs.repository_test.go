package pg_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repository/pg"
	"github.com/RafaLopesMelo/go-video-encoder/test"
	"github.com/stretchr/testify/require"
)

var videoId *vo.UniqueEntityID

func jobTestSetup() {
	connection := pg.NewConnection()
	repo := pg.NewVideosRepository(connection)

	video := test.DummyVideo()
	repo.Save(video)
	videoId = video.Video().ID
}

func TestCreateJob(t *testing.T) {
	connection := pg.NewConnection()
	repo := pg.NewJobsRepository(connection)

	dummy := test.DummyJob(videoId)

	repo.Save(dummy)
	job, err := repo.FindByID(*dummy.Job().ID)

	require.Nil(t, err)
	require.Equal(t, dummy.Job().ID, job.ID)
}

func TestUpdateJob(t *testing.T) {
	connection := pg.NewConnection()
	repo := pg.NewJobsRepository(connection)

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
