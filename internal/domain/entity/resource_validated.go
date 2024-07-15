package entity

type ValidatedResource struct {
	resource Resource
}

func (vr ValidatedResource) Resource() Resource {
	return vr.resource
}

func NewValidatedResource(resource Resource) (*ValidatedResource, error) {
	err := resource.validate()

	if err != nil {
		return nil, err
	}

	return &ValidatedResource{
		resource: resource,
	}, nil
}
