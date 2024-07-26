package pg_test

import (
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repository/pg"
	"github.com/RafaLopesMelo/go-video-encoder/test"
	"github.com/stretchr/testify/require"
)

func TestCreateJob(t *testing.T) {
	connection := pg.NewConnection()
	repo := pg.NewJobsRepository(connection)

	dummy := test.DummyJob(videoId)

	err := repo.Save(dummy)
	require.Nil(t, err)

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
	job.Error = "My error"
	validated, _ := entity.NewValidatedJob(job)

	repo.Save(validated)
	updated, err := repo.FindByID(*job.ID)

	require.Nil(t, err)
	require.Equal(t, job.ID, updated.ID)
	require.Equal(t, job.Error, updated.Error)
}
