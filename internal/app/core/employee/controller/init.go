package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service"
	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
	"github.com/rodericusifo/fiber-template/internal/pkg/middleware"
)

type EmployeeController struct {
	EmployeeService service.IEmployeeService
}

func InitEmployeeController(employeeService service.IEmployeeService) *EmployeeController {
	return &EmployeeController{EmployeeService: employeeService}
}

func (employeeController *EmployeeController) Mount(group fiber.Router) {
	group.Post("/create", middleware.HTTPUserRequest(), middleware.HTTPUserRolesPermission(constant.ADMIN), employeeController.CreateEmployee)
	group.Get("/list", middleware.HTTPUserRequest(), middleware.HTTPUserRolesPermission(constant.ADMIN), employeeController.GetListEmployee)
	group.Get("/:xid/detail", middleware.HTTPUserRequest(), middleware.HTTPUserRolesPermission(constant.ADMIN), employeeController.GetEmployee)
	group.Patch("/:xid/update", middleware.HTTPUserRequest(), middleware.HTTPUserRolesPermission(constant.ADMIN), employeeController.UpdateEmployee)
	group.Delete("/:xid/delete", middleware.HTTPUserRequest(), middleware.HTTPUserRolesPermission(constant.ADMIN), employeeController.DeleteEmployee)
}
