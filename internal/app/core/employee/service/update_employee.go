package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/input"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func (s *EmployeeService) UpdateEmployee(payload *input.UpdateEmployeeDTO) error {
	employeeModelRes, err := s.EmployeeResource.FirstEmployee(&pkg_types.QuerySQL{
		Searches: [][]pkg_types.SearchQuerySQLOperation{
			{
				{Field: "xid", Operator: "=", Value: payload.XID},
				{Field: "user_id", Operator: "=", Value: payload.UserID},
			},
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, "employee not found")
		}
		return err
	}

	employeeModel := employeeModelRes

	employeeModel.Address = payload.Address
	employeeModel.Age = payload.Age
	employeeModel.Birthday = payload.Birthday

	err = s.EmployeeResource.SaveEmployee(employeeModel)
	if err != nil {
		return err
	}

	return nil
}
