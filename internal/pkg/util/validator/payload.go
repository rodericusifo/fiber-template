package validator

import (
	"github.com/rodericusifo/fiber-template/internal/pkg/types"

	log "github.com/sirupsen/logrus"
)

type IPayload interface {
	CustomValidatePayload() error
}

func ValidatePayload(payload IPayload) error {
	validator := types.InitValidator()

	if err := validator.Validate(payload); err != nil {
		log.WithFields(log.Fields{
			"message": "validate payload fail",
			"detail":  err,
		}).Errorln("[VALIDATE PAYLOAD]")
		return err
	}
	if err := payload.CustomValidatePayload(); err != nil {
		return err
	}
	return nil
}
