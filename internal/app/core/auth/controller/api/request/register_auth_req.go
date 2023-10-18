package request

type RegisterAuthRequestBody struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	// RoleSlug string `json:"role_slug" validate:"required"`
}

func (r *RegisterAuthRequestBody) CustomValidateRequestBody() error {
	return nil
}
