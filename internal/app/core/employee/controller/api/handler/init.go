package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service"
	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
	"github.com/rodericusifo/fiber-template/internal/pkg/middleware"
)

type EmployeeHandler struct {
	EmployeeService service.IEmployeeService
}

func InitEmployeeHandler(employeeService service.IEmployeeService) *EmployeeHandler {
	return &EmployeeHandler{EmployeeService: employeeService}
}

func (employeeHandler *EmployeeHandler) Mount(group fiber.Router) {
	group.Post("/create", middleware.HTTPUserRequest(), middleware.HTTPUserRolesPermission(constant.ADMIN), employeeHandler.CreateEmployee)
	group.Get("/list", middleware.HTTPUserRequest(), middleware.HTTPUserRolesPermission(constant.ADMIN), employeeHandler.GetEmployees)
	group.Get("/:xid/detail", middleware.HTTPUserRequest(), middleware.HTTPUserRolesPermission(constant.ADMIN), employeeHandler.GetEmployee)
	group.Patch("/:xid/update", middleware.HTTPUserRequest(), middleware.HTTPUserRolesPermission(constant.ADMIN), employeeHandler.UpdateEmployee)
	group.Delete("/:xid/delete", middleware.HTTPUserRequest(), middleware.HTTPUserRolesPermission(constant.ADMIN), employeeHandler.DeleteEmployee)
}
