package service

import (
	"errors"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/output"
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
	pkg_util_counter "github.com/rodericusifo/fiber-template/pkg/util/counter"
	pkg_util_definer "github.com/rodericusifo/fiber-template/pkg/util/definer"
)

func init() {
	SetupTestEmployeeService()
}

func TestEmployeeService_GetEmployees(t *testing.T) {
	type (
		args struct {
			payload *input.GetEmployeesDTO
		}
		result struct {
			value output.GetEmployeesDTO
			meta  *pkg_types.Meta
			err   error
		}
	)

	testCases := []struct {
		desc   string
		input  args
		output result
		before func()
		after  func()
	}{
		{
			desc: "[ERROR]_because_something_error_happens",
			input: args{
				payload: &input.GetEmployeesDTO{
					Page:   &mockPage,
					Limit:  &mockLimit,
					UserID: 1,
				},
			},
			output: result{
				value: nil,
				meta:  nil,
				err:   errors.New("error something"),
			},
			before: func() {
				{
					var (
						page, limit = pkg_util_definer.DefinePaginationPageLimit(&mockPage, &mockLimit)
						offset      = pkg_util_counter.CountPaginationOffset(page, limit)

						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Limit:  limit,
							Offset: offset,
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result []*sql.Employee = nil
						err    error           = errors.New("error something")
					)
					mockEmployeeResource.EXPECT().FindEmployees(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_employees_not_found",
			input: args{
				payload: &input.GetEmployeesDTO{
					Page:   &mockPage,
					Limit:  &mockLimit,
					UserID: 1,
				},
			},
			output: result{
				value: nil,
				meta:  nil,
				err:   fiber.NewError(fiber.StatusNotFound, "employees not found"),
			},
			before: func() {
				{
					var (
						page, limit = pkg_util_definer.DefinePaginationPageLimit(&mockPage, &mockLimit)
						offset      = pkg_util_counter.CountPaginationOffset(page, limit)

						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Limit:  limit,
							Offset: offset,
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result []*sql.Employee = []*sql.Employee{}
						err    error           = nil
					)
					mockEmployeeResource.EXPECT().FindEmployees(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_error_on_count_all_employee",
			input: args{
				payload: &input.GetEmployeesDTO{
					Page:   &mockPage,
					Limit:  &mockLimit,
					UserID: 1,
				},
			},
			output: result{
				value: nil,
				meta:  nil,
				err:   errors.New("error count all employee"),
			},
			before: func() {
				{
					var (
						page, limit = pkg_util_definer.DefinePaginationPageLimit(&mockPage, &mockLimit)
						offset      = pkg_util_counter.CountPaginationOffset(page, limit)

						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Limit:  limit,
							Offset: offset,
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result []*sql.Employee = []*sql.Employee{
							{
								ID:        1,
								XID:       mockUUID,
								Name:      "Someone",
								Email:     "someone@mail.com",
								Address:   nil,
								Age:       nil,
								Birthday:  nil,
								CreatedAt: mockDate,
								UpdatedAt: mockDate,
							},
						}
						err error = nil
					)
					mockEmployeeResource.EXPECT().FindEmployees(arg1).Return(result, err).Once()
				}
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "user_id", Operator: "=", Value: uint(1)},
									{Field: "deleted_at", Operator: "IS NULL"},
								},
							},
						}
					)
					var (
						count int64 = 0
						err   error = errors.New("error count all employee")
					)
					mockEmployeeResource.EXPECT().CountEmployees(arg1).Return(count, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_get_employees",
			input: args{
				payload: &input.GetEmployeesDTO{
					Page:   &mockPage,
					Limit:  &mockLimit,
					UserID: 1,
				},
			},
			output: result{
				value: []*output.EmployeeDTO{
					{
						XID:       mockUUID,
						Name:      "Someone",
						Email:     "someone@mail.com",
						Address:   nil,
						Age:       nil,
						Birthday:  nil,
						CreatedAt: mockDate,
						UpdatedAt: mockDate,
					},
				},
				meta: &pkg_types.Meta{
					CurrentPage:      1,
					CountDataPerPage: 1,
					TotalData:        1,
					TotalPage:        1,
				},
				err: nil,
			},
			before: func() {
				{
					var (
						page, limit = pkg_util_definer.DefinePaginationPageLimit(&mockPage, &mockLimit)
						offset      = pkg_util_counter.CountPaginationOffset(page, limit)

						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Limit:  limit,
							Offset: offset,
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result []*sql.Employee = []*sql.Employee{
							{
								ID:        1,
								XID:       mockUUID,
								Name:      "Someone",
								Email:     "someone@mail.com",
								Address:   nil,
								Age:       nil,
								Birthday:  nil,
								CreatedAt: mockDate,
								UpdatedAt: mockDate,
							},
						}
						err error = nil
					)
					mockEmployeeResource.EXPECT().FindEmployees(arg1).Return(result, err).Once()
				}
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "user_id", Operator: "=", Value: uint(1)},
									{Field: "deleted_at", Operator: "IS NULL"},
								},
							},
						}
					)
					var (
						count int64 = 1
						err   error = nil
					)
					mockEmployeeResource.EXPECT().CountEmployees(arg1).Return(count, err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, meta, err := employeeService.GetEmployees(tC.input.payload)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.meta, meta)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
