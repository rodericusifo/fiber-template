package validator

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/pkg/types"

	log "github.com/sirupsen/logrus"
)

type IRequestBody interface {
	CustomValidateRequestBody() error
}

func ValidateRequestBody(ctx *fiber.Ctx, req IRequestBody) error {
	validator := types.InitValidator()

	if err := ctx.BodyParser(req); err != nil {
		log.WithFields(log.Fields{
			"message": "bind request body fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST BODY]")
		return err
	}
	if err := validator.Validate(req); err != nil {
		log.WithFields(log.Fields{
			"message": "validate request body fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST BODY]")
		return err
	}
	if err := req.CustomValidateRequestBody(); err != nil {
		return err
	}
	return nil
}
