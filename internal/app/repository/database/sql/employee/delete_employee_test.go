package employee

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
)

func init() {
	SetupTestMysqlEmployeeDatabaseSQLRepository()
}

func TestMysqlEmployeeDatabaseSQLRepository_DeleteEmployee(t *testing.T) {
	type (
		args struct {
			payload *sql.Employee
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
			desc: "[ERROR]_because_something_error_happens",
			input: args{
				payload: &sql.Employee{
					ID: 1,
				},
			},
			output: result{
				err: errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 = 1
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							"DELETE FROM `employees` WHERE `employees`.`id` = ?",
						),
					).
						WithArgs(arg1).
						WillReturnError(errors.New("error something"))
					mockQuery.ExpectRollback()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_delete_employee",
			input: args{
				payload: &sql.Employee{
					ID: 1,
				},
			},
			output: result{
				err: nil,
			},
			before: func() {
				{
					var (
						arg1 = 1
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							"DELETE FROM `employees` WHERE `employees`.`id` = ?",
						),
					).
						WithArgs(arg1).
						WillReturnResult(sqlmock.NewResult(0, 1))
					mockQuery.ExpectCommit()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			err := employeeDatabaseSQLRepository.DeleteEmployee(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
