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

func TestEmployeeService_CreateEmployee(t *testing.T) {
	type (
		args struct {
			payload *input.CreateEmployeeDTO
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
			desc: "[ERROR]_because_error_something_on_get_employee",
			input: args{
				payload: &input.CreateEmployeeDTO{
					Name:   "Someone",
					Email:  "someone@mail.com",
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
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
								},
							},
							WithDeleted: true,
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
			desc: "[ERROR]_because_employee_already_registered",
			input: args{
				payload: &input.CreateEmployeeDTO{
					Name:   "Someone",
					Email:  "someone@mail.com",
					UserID: 1,
				},
			},
			output: result{
				err: fiber.NewError(fiber.StatusConflict, "employee already registered"),
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
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
								},
							},
							WithDeleted: true,
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
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_error_something_on_create_employee",
			input: args{
				payload: &input.CreateEmployeeDTO{
					Name:   "Someone",
					Email:  "someone@mail.com",
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
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
								},
							},
							WithDeleted: true,
						}
					)
					var (
						result *sql.Employee = nil
						err    error         = gorm.ErrRecordNotFound
					)
					mockEmployeeResource.EXPECT().FirstEmployee(arg1).Return(result, err).Once()
				}
				{
					var (
						arg1 *sql.Employee = &sql.Employee{
							Name:   "Someone",
							Email:  "someone@mail.com",
							UserID: 1,
						}
					)
					var (
						err error = errors.New("error something")
					)
					mockEmployeeResource.EXPECT().SaveEmployee(arg1).Return(err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_create_employee",
			input: args{
				payload: &input.CreateEmployeeDTO{
					Name:   "Someone",
					Email:  "someone@mail.com",
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
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
								},
							},
							WithDeleted: true,
						}
					)
					var (
						result *sql.Employee = nil
						err    error         = gorm.ErrRecordNotFound
					)
					mockEmployeeResource.EXPECT().FirstEmployee(arg1).Return(result, err).Once()
				}
				{
					var (
						arg1 *sql.Employee = &sql.Employee{
							Name:   "Someone",
							Email:  "someone@mail.com",
							UserID: 1,
						}
					)
					var (
						err error = nil
					)
					mockEmployeeResource.EXPECT().SaveEmployee(arg1).Return(err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			err := employeeService.CreateEmployee(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
