package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/app/core/employee/controller/request"
	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/getter"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/serializer"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/validator"

	pkg_util_response "github.com/rodericusifo/fiber-template/pkg/util/response"
)

func (c *EmployeeController) GetEmployee(ctx *fiber.Ctx) error {
	reqUser := getter.GetRequestUser(ctx)

	reqParams := new(request.GetEmployeeRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	employeeDtoRes, err := c.EmployeeService.GetEmployee(&input.GetEmployeeDTO{
		XID:    reqParams.XID,
		UserID: reqUser.ID,
	})
	if err != nil {
		return err
	}

	getEmployeeResponse := serializer.SerializeEmployeeDTOToEmployeeResponse(employeeDtoRes)

	return ctx.Status(fiber.StatusOK).JSON(pkg_util_response.ResponseSuccess("get employee success", getEmployeeResponse, nil))
}
