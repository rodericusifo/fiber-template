package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/pkg/util/handler"

	lib_mockery_mocks "github.com/rodericusifo/fiber-template/lib/mockery/mocks"
)

var (
	mockApp         *fiber.App
	mockAuthService *lib_mockery_mocks.IAuthService
	authController  *AuthController
)

var (
	mockEmail, mockPassword, mockJWTToken string
)

func SetupTestAuthController() {
	mockApp = fiber.New(fiber.Config{
		ErrorHandler: handler.HandleHTTPError,
	})

	mockAuthService = new(lib_mockery_mocks.IAuthService)

	auth := mockApp.Group("/auth")
	authController = InitAuthController(mockAuthService)
	authController.Mount(auth)

	mockEmail = "john@gmail.com"
	mockPassword = "john1223"

	mockJWTToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjhlYTc3OGJjLTM5NTgtNGU5Zi04ZmEyLWE4YTlhZDhmMmFiMSIsIm5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AZ21haWwuY29tIiwicm9sZSI6IkFETUlOIiwiZXhwIjoxNjc3MDc5NzgxfQ.bndXk_BggjadIF2Rwluxc-3tPr-ArfWVYTZ5y03wHU8"
}
