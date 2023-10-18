package request

type UpdateEmployeeRequestBody struct {
	Address  *string `json:"address" validate:"omitempty"`
	Age      *int    `json:"age" validate:"omitempty,min=1"`
	Birthday *string `json:"birthday" validate:"omitempty"`
}

func (r *UpdateEmployeeRequestBody) CustomValidateRequestBody() error {
	return nil
}

type UpdateEmployeeRequestParams struct {
	XID string `param:"xid" validate:"required,uuid4"`
}

func (r *UpdateEmployeeRequestParams) CustomValidateRequestParams() error {
	return nil
}
