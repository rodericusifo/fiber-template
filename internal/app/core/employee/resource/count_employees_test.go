package resource

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func init() {
	SetupTestEmployeeResource()
}

func TestEmployeeResource_CountEmployees(t *testing.T) {
	type (
		args struct {
			query *pkg_types.QuerySQL
		}
		result struct {
			value int64
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
					Searches: [][]pkg_types.SearchQuerySQLOperation{
						{
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
				},
			},
			output: result{
				value: 0,
				err:   errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result int64 = 0
						err    error = errors.New("error something")
					)
					mockEmployeeDatabaseSQLRepository.EXPECT().CountEmployees(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_count_all_employee",
			input: args{
				query: &pkg_types.QuerySQL{
					Searches: [][]pkg_types.SearchQuerySQLOperation{
						{
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
				},
			},
			output: result{
				value: 1,
				err:   nil,
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result int64 = 1
						err    error = nil
					)
					mockEmployeeDatabaseSQLRepository.EXPECT().CountEmployees(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, err := employeeResource.CountEmployees(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
