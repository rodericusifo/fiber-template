package validator

import (
	"github.com/rodericusifo/fiber-template/internal/pkg/types"

	log "github.com/sirupsen/logrus"
)

type IRequestUser interface {
	CustomValidateRequestUser() error
}

func ValidateRequestUser(req IRequestUser) error {
	validator := types.InitValidator()

	if err := validator.Validate(req); err != nil {
		log.WithFields(log.Fields{
			"message": "validate request user fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST USER]")
		return err
	}
	if err := req.CustomValidateRequestUser(); err != nil {
		return err
	}
	return nil
}
