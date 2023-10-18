package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/app/core/auth/service"
)

type AuthController struct {
	AuthService service.IAuthService
}

func InitAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (authController *AuthController) Mount(group fiber.Router) {
	group.Post("/register", authController.RegisterAuth)
	group.Post("/login", authController.LoginAuth)
}
