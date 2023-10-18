package types

type RequestRole struct {
	ID   uint   `validate:"required"`
	Slug string `validate:"required"`
}

type RequestUser struct {
	ID    uint   `validate:"required"`
	XID   string `validate:"required,uuid4"`
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Role  RequestRole
}

func (r *RequestUser) CustomValidateRequestUser() error {
	return nil
}
