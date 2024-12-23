package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func (s *EmployeeService) CreateEmployee(payload *input.CreateEmployeeDTO) error {
	employeeModelRes, err := s.EmployeeResource.FirstEmployee(&pkg_types.QuerySQL{
		Selects: []pkg_types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]pkg_types.SearchQuerySQLOperation{
			{
				{Field: "email", Operator: "=", Value: payload.Email},
			},
		},
		WithDeleted: true,
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if employeeModelRes != nil {
		return fiber.NewError(fiber.StatusConflict, "employee already registered")
	}

	employeeModel := &sql.Employee{
		Name:     payload.Name,
		Email:    payload.Email,
		Address:  payload.Address,
		Age:      payload.Age,
		Birthday: payload.Birthday,
		UserID:   payload.UserID,
	}
	err = s.EmployeeResource.SaveEmployee(employeeModel)
	if err != nil {
		return err
	}

	return nil
}
