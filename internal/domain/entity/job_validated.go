package entity

type ValidatedJob struct {
	job Job
}

func (vj ValidatedJob) Job() Job {
	return vj.job
}

func NewValidatedJob(job Job) (ValidatedJob, error) {
	err := job.validate()

	if err != nil {
		return ValidatedJob{}, err
	}

	return ValidatedJob{
		job: job,
	}, nil
}
