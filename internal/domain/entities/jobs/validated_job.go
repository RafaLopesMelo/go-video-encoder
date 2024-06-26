package jobs

type ValidatedJob struct {
    job Job
}

func NewValidatedJob(job Job) (*ValidatedJob, error) {
	err := job.validate()

	if err != nil {
		return nil, err
	}

	return &ValidatedJob{
		job: job,
	}, nil
}
