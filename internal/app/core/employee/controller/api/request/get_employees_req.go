package request

type GetEmployeesRequestQuery struct {
	Page  *int `query:"page" validate:"omitempty,min=0"`
	Limit *int `query:"limit" validate:"omitempty,min=0"`
}

func (r *GetEmployeesRequestQuery) CustomValidateRequestQuery() error {
	return nil
}
