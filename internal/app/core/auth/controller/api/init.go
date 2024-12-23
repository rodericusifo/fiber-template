package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/app/core/auth/controller/api/handler"

	lib_wire_core_service_auth "github.com/rodericusifo/fiber-template/lib/wire/core/service/auth"
)

func InitAPI(router fiber.Router) {
	auth := router.Group("/auth")
	authService := lib_wire_core_service_auth.AuthService()
	authHandler := handler.InitAuthHandler(authService)
	authHandler.Mount(auth)
}
