package types

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator *validator.Validate
}

func InitValidator() *Validator {
	return &Validator{Validator: validator.New()}
}

func (cv *Validator) Validate(i any) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}
