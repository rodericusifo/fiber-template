package handler

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	log "github.com/sirupsen/logrus"

	pkg_util_response "github.com/rodericusifo/fiber-template/pkg/util/response"
)

func HandleHTTPError(ctx *fiber.Ctx, err error) error {
	log.WithFields(log.Fields{
		"type":   fmt.Sprintf("%T", err),
		"detail": err,
	}).Errorln("[HANDLE HTTP ERROR]")
	fe, ok := err.(*fiber.Error)
	if ok {
		return ctx.Status(fe.Code).JSON(pkg_util_response.ResponseFail(fmt.Sprint(fe.Error()), fe))
	}
	ve, ok := err.(validator.ValidationErrors)
	if ok {
		type ErrorResponse struct {
			FailedField string `json:"failed_field"`
			Tag         string `json:"tag"`
			Error       string `json:"error"`
		}
		var errors []*ErrorResponse
		for _, err := range ve {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Error = err.Error()
			errors = append(errors, &element)
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_util_response.ResponseFail("validation error", errors))
	}
	me, ok := err.(*json.MarshalerError)
	if ok {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(pkg_util_response.ResponseFail(me.Error(), me.Unwrap()))
	}
	re, ok := err.(runtime.Error)
	if ok {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(pkg_util_response.ResponseFail(re.Error(), re))
	}
	if e := err; e != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(pkg_util_response.ResponseFail(e.Error(), e))
	}
	return nil
}
