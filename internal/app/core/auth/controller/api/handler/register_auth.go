package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/app/core/auth/controller/api/request"
	"github.com/rodericusifo/fiber-template/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/validator"

	pkg_util_response "github.com/rodericusifo/fiber-template/pkg/util/response"
)

func (h *AuthHandler) RegisterAuth(ctx *fiber.Ctx) error {
	reqBody := new(request.RegisterAuthRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	err := h.AuthService.RegisterAuth(&input.RegisterAuthDTO{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Password: reqBody.Password,
		// RoleSlug: reqBody.RoleSlug,
	})
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(pkg_util_response.ResponseSuccess[any]("auth register success", nil, nil))
}
