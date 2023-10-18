package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/app/core/employee/controller/api/request"
	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/getter"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/validator"

	pkg_util_response "github.com/rodericusifo/fiber-template/pkg/util/response"
)

func (h *EmployeeHandler) UpdateEmployee(ctx *fiber.Ctx) error {
	reqUser := getter.GetRequestUser(ctx)

	reqBody := new(request.UpdateEmployeeRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	reqParams := new(request.UpdateEmployeeRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	birthdayString := *reqBody.Birthday
	birthdayTime, err := time.Parse(constant.DEFAULT_TIME_LAYOUT.(string), birthdayString)
	if err != nil {
		return err
	}

	if err := h.EmployeeService.UpdateEmployee(&input.UpdateEmployeeDTO{
		XID:      reqParams.XID,
		Address:  reqBody.Address,
		Age:      reqBody.Age,
		Birthday: &birthdayTime,
		UserID:   reqUser.ID,
	}); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(pkg_util_response.ResponseSuccess[any]("update employee success", nil, nil))
}
