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

func (c *EmployeeController) GetListEmployee(ctx *fiber.Ctx) error {
	reqUser := getter.GetRequestUser(ctx)

	reqQuery := new(request.GetListEmployeeRequestQuery)
	if err := validator.ValidateRequestQuery(ctx, reqQuery); err != nil {
		return err
	}

	employeeListDtoRes, meta, err := c.EmployeeService.GetListEmployee(&input.GetListEmployeeDTO{
		Page:   reqQuery.Page,
		Limit:  reqQuery.Limit,
		UserID: reqUser.ID,
	})
	if err != nil {
		return err
	}

	getListEmployeeResponse := serializer.SerializeEmployeeDTOsToEmployeeResponses(employeeListDtoRes)

	return ctx.Status(fiber.StatusOK).JSON(pkg_util_response.ResponseSuccess("get list employee success", getListEmployeeResponse, meta))
}
