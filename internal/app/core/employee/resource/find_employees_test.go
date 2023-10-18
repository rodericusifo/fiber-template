package resource

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
	pkg_util_counter "github.com/rodericusifo/fiber-template/pkg/util/counter"
)

func init() {
	SetupTestEmployeeResource()
}

func TestEmployeeResource_FindEmployees(t *testing.T) {
	type (
		args struct {
			query *pkg_types.QuerySQL
		}
		result struct {
			value []*sql.Employee
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
				query: &pkg_types.QuerySQL{
					Offset: pkg_util_counter.CountPaginationOffset(1, 10),
					Limit:  10,
					Searches: [][]pkg_types.SearchQuerySQLOperation{
						{
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
				},
			},
			output: result{
				value: nil,
				err:   errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Offset: pkg_util_counter.CountPaginationOffset(1, 10),
							Limit:  10,
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
					mockEmployeeDatabaseSQLRepository.EXPECT().FindEmployees(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_find_employees",
			input: args{
				query: &pkg_types.QuerySQL{
					Offset: pkg_util_counter.CountPaginationOffset(1, 10),
					Limit:  10,
					Searches: [][]pkg_types.SearchQuerySQLOperation{
						{
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
				},
			},
			output: result{
				value: []*sql.Employee{
					{
						ID:        1,
						XID:       mockUUID,
						Name:      "Ifo",
						Email:     "Ifo@gmail.com",
						Address:   &mockAddress,
						Age:       &mockAge,
						Birthday:  &mockBirthday,
						UserID:    1,
						CreatedAt: mockDateTime,
						UpdatedAt: mockDateTime,
					},
				},
				err: nil,
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Offset: pkg_util_counter.CountPaginationOffset(1, 10),
							Limit:  10,
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
								Name:      "Ifo",
								Email:     "Ifo@gmail.com",
								Address:   &mockAddress,
								Age:       &mockAge,
								Birthday:  &mockBirthday,
								UserID:    1,
								CreatedAt: mockDateTime,
								UpdatedAt: mockDateTime,
							},
						}
						err error = nil
					)
					mockEmployeeDatabaseSQLRepository.EXPECT().FindEmployees(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, err := employeeResource.FindEmployees(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
