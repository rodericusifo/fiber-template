package validator

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/pkg/types"

	log "github.com/sirupsen/logrus"
)

type IRequestQuery interface {
	CustomValidateRequestQuery() error
}

func ValidateRequestQuery(ctx *fiber.Ctx, req IRequestQuery) error {
	validator := types.InitValidator()

	if err := ctx.QueryParser(req); err != nil {
		log.WithFields(log.Fields{
			"message": "bind request query fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST QUERY]")
		return err
	}
	if err := validator.Validate(req); err != nil {
		log.WithFields(log.Fields{
			"message": "validate request query fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST QUERY]")
		return err
	}
	if err := req.CustomValidateRequestQuery(); err != nil {
		return err
	}
	return nil
}
