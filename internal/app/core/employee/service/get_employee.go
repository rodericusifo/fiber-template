package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/output"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/serializer"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func (s *EmployeeService) GetEmployee(payload *input.GetEmployeeDTO) (output.GetEmployeeDTO, error) {
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
			return nil, fiber.NewError(fiber.StatusNotFound, "employee not found")
		}
		return nil, err
	}

	employeeDto := serializer.SerializeEmployeeToEmployeeDTO(employeeModelRes)

	return employeeDto, nil
}
