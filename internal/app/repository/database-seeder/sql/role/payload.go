package role

type RoleSeedPayload struct {
	XID  string `validate:"required,uuid4"`
	Name string `validate:"required"`
	Slug string `validate:"required"`
}

func (r *RoleSeedPayload) CustomValidatePayload() error {
	return nil
}
