package employee

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func init() {
	SetupTestMysqlEmployeeDatabaseSQLRepository()
}

func TestMysqlEmployeeDatabaseSQLRepository_CountEmployees(t *testing.T) {
	type (
		args struct {
			query *pkg_types.QuerySQL
		}
		result struct {
			count int64
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
				query: nil,
			},
			output: result{
				count: 0,
				err:   errors.New("something error"),
			},
			before: func() {
				{
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT COUNT(`employees`.`id`) FROM `employees`",
						),
					).
						WillReturnError(errors.New("something error"))
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_count_employees",
			input: args{
				query: &pkg_types.QuerySQL{
					Searches: [][]pkg_types.SearchQuerySQLOperation{
						{
							{Field: "user_id", Operator: "=", Value: uint(1)},
							{Field: "deleted_at", Operator: "IS NULL"},
						},
					},
				},
			},
			output: result{
				count: 1,
				err:   nil,
			},
			before: func() {
				{
					var (
						arg1         = 1
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT COUNT(`employees`.`id`) FROM `employees` WHERE `employees`.`user_id` = ? AND `employees`.`deleted_at` IS NULL",
						),
					).
						WithArgs(arg1).
						WillReturnRows(rowsInstance)
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			count, err := employeeDatabaseSQLRepository.CountEmployees(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.count, count)

			tC.after()
		})
	}
}
