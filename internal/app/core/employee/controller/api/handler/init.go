package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service"
	"github.com/rodericusifo/fiber-template/internal/pkg/middleware"
)

type EmployeeHandler struct {
	EmployeeService service.IEmployeeService
}

func InitEmployeeHandler(employeeService service.IEmployeeService) *EmployeeHandler {
	return &EmployeeHandler{EmployeeService: employeeService}
}

func (employeeHandler *EmployeeHandler) Mount(group fiber.Router) {
	group.Post("/create", middleware.HTTPUserRequest(), middleware.HTTPUserRolePermissions(), employeeHandler.CreateEmployee)
	group.Get("/list", middleware.HTTPUserRequest(), middleware.HTTPUserRolePermissions(), employeeHandler.GetEmployees)
	group.Get("/:xid/detail", middleware.HTTPUserRequest(), middleware.HTTPUserRolePermissions(), employeeHandler.GetEmployee)
	group.Put("/:xid/update", middleware.HTTPUserRequest(), middleware.HTTPUserRolePermissions(), employeeHandler.UpdateEmployee)
	group.Delete("/:xid/delete", middleware.HTTPUserRequest(), middleware.HTTPUserRolePermissions(), employeeHandler.DeleteEmployee)
}
