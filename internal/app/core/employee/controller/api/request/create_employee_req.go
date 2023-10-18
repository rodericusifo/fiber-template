package request

type CreateEmployeeRequestBody struct {
	Name     string  `json:"name" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Address  *string `json:"address" validate:"omitempty"`
	Age      *int    `json:"age" validate:"omitempty,min=0"`
	Birthday *string `json:"birthday" validate:"omitempty"`
}

func (r *CreateEmployeeRequestBody) CustomValidateRequestBody() error {
	return nil
}
