package service

import (
	"errors"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func init() {
	SetupTestEmployeeService()
}

func TestEmployeeService_DeleteEmployee(t *testing.T) {
	type (
		args struct {
			payload *input.DeleteEmployeeDTO
		}
		result struct {
			err error
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
			desc: "[ERROR]_because_employee_not_found",
			input: args{
				payload: &input.DeleteEmployeeDTO{
					XID:    mockUUID,
					UserID: 1,
				},
			},
			output: result{
				err: fiber.NewError(fiber.StatusNotFound, "employee not found"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Selects: []pkg_types.SelectQuerySQLOperation{
								{Field: "id"},
							},
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "xid", Operator: "=", Value: mockUUID},
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result *sql.Employee = nil
						err    error         = gorm.ErrRecordNotFound
					)
					mockEmployeeResource.EXPECT().FirstEmployee(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_something_error_happens",
			input: args{
				payload: &input.DeleteEmployeeDTO{
					XID:    mockUUID,
					UserID: 1,
				},
			},
			output: result{
				err: errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Selects: []pkg_types.SelectQuerySQLOperation{
								{Field: "id"},
							},
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "xid", Operator: "=", Value: mockUUID},
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result *sql.Employee = nil
						err    error         = errors.New("error something")
					)
					mockEmployeeResource.EXPECT().FirstEmployee(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_error_on_delete",
			input: args{
				payload: &input.DeleteEmployeeDTO{
					XID:    mockUUID,
					UserID: 1,
				},
			},
			output: result{
				err: errors.New("error delete"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Selects: []pkg_types.SelectQuerySQLOperation{
								{Field: "id"},
							},
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "xid", Operator: "=", Value: mockUUID},
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result *sql.Employee = &sql.Employee{
							ID: 1,
						}
						err error = nil
					)
					mockEmployeeResource.EXPECT().FirstEmployee(arg1).Return(result, err).Once()
				}
				{
					var (
						arg1 *sql.Employee = &sql.Employee{
							ID: 1,
						}
					)
					var (
						err error = errors.New("error delete")
					)
					mockEmployeeResource.EXPECT().DeleteEmployee(arg1).Return(err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_delete_employee",
			input: args{
				payload: &input.DeleteEmployeeDTO{
					XID:    mockUUID,
					UserID: 1,
				},
			},
			output: result{
				err: nil,
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Selects: []pkg_types.SelectQuerySQLOperation{
								{Field: "id"},
							},
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "xid", Operator: "=", Value: mockUUID},
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result *sql.Employee = &sql.Employee{
							ID: 1,
						}
						err error = nil
					)
					mockEmployeeResource.EXPECT().FirstEmployee(arg1).Return(result, err).Once()
				}
				{
					var (
						arg1 *sql.Employee = &sql.Employee{
							ID: 1,
						}
					)
					var (
						err error = nil
					)
					mockEmployeeResource.EXPECT().DeleteEmployee(arg1).Return(err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			err := employeeService.DeleteEmployee(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
