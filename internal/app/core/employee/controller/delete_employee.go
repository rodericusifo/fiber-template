package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/app/core/employee/controller/request"
	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/getter"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/validator"

	pkg_util_response "github.com/rodericusifo/fiber-template/pkg/util/response"
)

func (c *EmployeeController) DeleteEmployee(ctx *fiber.Ctx) error {
	reqUser := getter.GetRequestUser(ctx)

	reqParams := new(request.DeleteEmployeeRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	if err := c.EmployeeService.DeleteEmployee(&input.DeleteEmployeeDTO{
		XID:    reqParams.XID,
		UserID: reqUser.ID,
	}); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(pkg_util_response.ResponseSuccess[any]("delete employee success", nil, nil))
}
