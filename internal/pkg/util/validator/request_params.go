package validator

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/pkg/types"

	log "github.com/sirupsen/logrus"
)

type IRequestParams interface {
	CustomValidateRequestParams() error
}

func ValidateRequestParams(ctx *fiber.Ctx, req IRequestParams) error {
	validator := types.InitValidator()

	if err := ctx.ParamsParser(req); err != nil {
		log.WithFields(log.Fields{
			"message": "bind request params fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST PARAMS]")
		return err
	}
	if err := validator.Validate(req); err != nil {
		log.WithFields(log.Fields{
			"message": "validate request params fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST PARAMS]")
		return err
	}
	if err := req.CustomValidateRequestParams(); err != nil {
		return err
	}
	return nil
}
