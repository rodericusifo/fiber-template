package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/app/core/employee/controller/api/request"
	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/getter"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/serializer"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/validator"

	pkg_util_response "github.com/rodericusifo/fiber-template/pkg/util/response"
)

func (h *EmployeeHandler) GetEmployees(ctx *fiber.Ctx) error {
	reqUser := getter.GetRequestUser(ctx)

	reqQuery := new(request.GetEmployeesRequestQuery)
	if err := validator.ValidateRequestQuery(ctx, reqQuery); err != nil {
		return err
	}

	getEmployeesDtoRes, meta, err := h.EmployeeService.GetEmployees(&input.GetEmployeesDTO{
		Page:   reqQuery.Page,
		Limit:  reqQuery.Limit,
		UserID: reqUser.ID,
	})
	if err != nil {
		return err
	}

	getEmployeesResponse := serializer.SerializeEmployeeDTOsToEmployeeResponses(getEmployeesDtoRes)

	return ctx.Status(fiber.StatusOK).JSON(pkg_util_response.ResponseSuccess("get employees success", getEmployeesResponse, meta))
}
