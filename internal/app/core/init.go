package core

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/pkg/util/getter"

	jwtware "github.com/gofiber/contrib/jwt"

	internal_app_core_auth_controller "github.com/rodericusifo/fiber-template/internal/app/core/auth/controller"
	internal_app_core_employee_controller "github.com/rodericusifo/fiber-template/internal/app/core/employee/controller"
	lib_wire_core_service_auth "github.com/rodericusifo/fiber-template/lib/wire/core/service/auth"
	lib_wire_core_service_employee "github.com/rodericusifo/fiber-template/lib/wire/core/service/employee"
)

func InitRoutes(app *fiber.App) {
	{
		auth := app.Group("/auth")
		authService := lib_wire_core_service_auth.AuthService()
		authController := internal_app_core_auth_controller.InitAuthController(authService)
		authController.Mount(auth)
	}
	{
		employee := app.Group("/employees")
		employee.Use(jwtware.New(*getter.GetJWTAuthConfig()))
		employeeService := lib_wire_core_service_employee.EmployeeService()
		employeeController := internal_app_core_employee_controller.InitEmployeeController(employeeService)
		employeeController.Mount(employee)
	}
}
