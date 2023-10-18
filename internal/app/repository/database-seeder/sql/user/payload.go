package user

type UserSeedPayload struct {
	XID      string `validate:"required,uuid4"`
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func (r *UserSeedPayload) CustomValidatePayload() error {
	return nil
}
