package entity

type ValidatedResource struct {
	rw ResourceWrapper
}

func (vr ValidatedResource) Wrapper() ResourceWrapper {
	return vr.rw
}

func NewValidatedResource(rw ResourceWrapper) (*ValidatedResource, error) {
	err := rw.validate()

	if err != nil {
		return nil, err
	}

	resource := rw.Resource()
	err = resource.validate()

	if err != nil {
		return nil, err
	}

	return &ValidatedResource{
		rw: rw,
	}, nil
}
