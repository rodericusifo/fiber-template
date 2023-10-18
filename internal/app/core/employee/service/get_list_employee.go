package service

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/output"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/serializer"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
	pkg_util_counter "github.com/rodericusifo/fiber-template/pkg/util/counter"
	pkg_util_definer "github.com/rodericusifo/fiber-template/pkg/util/definer"
)

func (s *EmployeeService) GetListEmployee(payload *input.GetListEmployeeDTO) (output.GetListEmployeeDTO, *pkg_types.Meta, error) {
	page, limit := pkg_util_definer.DefinePaginationPageLimit(payload.Page, payload.Limit)

	employeeListModelRes, countEmployeeListModelRes, err := s.EmployeeResource.GetListEmployeeAndCount(&pkg_types.QuerySQL{
		Offset: pkg_util_counter.CountPaginationOffset(page, limit),
		Limit:  limit,
		Searches: [][]pkg_types.SearchQuerySQLOperation{
			{
				{Field: "user_id", Operator: "=", Value: payload.UserID},
			},
		},
	})
	if err != nil {
		return nil, nil, err
	}

	if len(employeeListModelRes) < 1 {
		return nil, nil, fiber.NewError(fiber.StatusNotFound, "list employee not found")
	}

	countEmployeeAllModelRes, err := s.EmployeeResource.CountAllEmployee(&pkg_types.QuerySQL{
		Searches: [][]pkg_types.SearchQuerySQLOperation{
			{
				{Field: "user_id", Operator: "=", Value: payload.UserID},
			},
		},
	})
	if err != nil {
		return nil, nil, err
	}

	employeeListDto := serializer.SerializeEmployeesToEmployeeDTOs(employeeListModelRes)

	meta := &pkg_types.Meta{
		CurrentPage:      page,
		CountDataPerPage: countEmployeeListModelRes,
		TotalData:        countEmployeeAllModelRes,
	}

	meta.TotalPage = pkg_util_counter.CountPaginationTotalPage(meta.CountDataPerPage, meta.TotalData)

	return employeeListDto, meta, nil
}
