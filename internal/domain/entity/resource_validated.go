package entity

type ValidatedResource struct {
	rw ResourceWrapper
}

func (vr ValidatedResource) Wrapper() ResourceWrapper {
	return vr.rw
}

func NewValidatedResource(rw ResourceWrapper) (ValidatedResource, error) {
	err := rw.validate()

	if err != nil {
		return ValidatedResource{}, err
	}

	r := rw.Resource()
	err = r.validate()

	if err != nil {
		return ValidatedResource{}, err
	}

	return ValidatedResource{
		rw: rw,
	}, nil
}
