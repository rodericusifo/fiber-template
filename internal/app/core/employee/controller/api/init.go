package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/app/core/employee/controller/api/handler"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/getter"

	jwtware "github.com/gofiber/contrib/jwt"

	lib_wire_core_service_employee "github.com/rodericusifo/fiber-template/lib/wire/core/service/employee"
)

func InitAPI(router fiber.Router) {
	employee := router.Group("/employees")
	employee.Use(jwtware.New(*getter.GetJWTAuthConfig()))
	employeeService := lib_wire_core_service_employee.EmployeeService()
	employeeHandler := handler.InitEmployeeHandler(employeeService)
	employeeHandler.Mount(employee)
}
