package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/app/core/auth/service"
)

type AuthHandler struct {
	AuthService service.IAuthService
}

func InitAuthHandler(authService service.IAuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (authHandler *AuthHandler) Mount(group fiber.Router) {
	group.Post("/register", authHandler.RegisterAuth)
	group.Post("/login", authHandler.LoginAuth)
}
