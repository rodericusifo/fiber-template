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

func (h *EmployeeHandler) CreateEmployee(ctx *fiber.Ctx) error {
	reqUser := getter.GetRequestUser(ctx)

	reqBody := new(request.CreateEmployeeRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	birthdayString := *reqBody.Birthday
	birthdayTime, err := time.Parse(constant.DEFAULT_TIME_LAYOUT.(string), birthdayString)
	if err != nil {
		return err
	}

	if err := h.EmployeeService.CreateEmployee(&input.CreateEmployeeDTO{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Address:  reqBody.Address,
		Age:      reqBody.Age,
		Birthday: &birthdayTime,
		UserID:   reqUser.ID,
	}); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(pkg_util_response.ResponseSuccess[any]("create employee success", nil, nil))
}
