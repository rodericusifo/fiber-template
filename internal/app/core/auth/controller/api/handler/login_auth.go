package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/app/core/auth/controller/api/request"
	"github.com/rodericusifo/fiber-template/internal/app/core/auth/controller/api/response"
	"github.com/rodericusifo/fiber-template/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/validator"

	pkg_util_response "github.com/rodericusifo/fiber-template/pkg/util/response"
)

func (h *AuthHandler) LoginAuth(ctx *fiber.Ctx) error {
	reqBody := new(request.LoginAuthRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	authLoginDtoRes, err := h.AuthService.LoginAuth(&input.LoginAuthDTO{
		Email:    reqBody.Email,
		Password: reqBody.Password,
	})
	if err != nil {
		return err
	}

	loginAuthRes := &response.LoginAuthResponse{
		Token: authLoginDtoRes.Token,
	}

	return ctx.Status(fiber.StatusOK).JSON(pkg_util_response.ResponseSuccess("auth login success", loginAuthRes, nil))
}
