package user

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func init() {
	SetupTestUserDatabaseSeederSQLRepository()
}

func TestUserDatabaseSeederSQLRepository_Clear(t *testing.T) {
	type (
		args struct {
			db *gorm.DB
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
				db: mockDB,
			},
			output: result{
				err: errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 = "super_admin"
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `roles`.`id` FROM `roles` WHERE `roles`.`slug` = ? ORDER BY `roles`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1).
						WillReturnError(errors.New("error something"))
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_something_error_happens#1",
			input: args{
				db: mockDB,
			},
			output: result{
				err: errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1         = "super_admin"
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `roles`.`id` FROM `roles` WHERE `roles`.`slug` = ? ORDER BY `roles`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1).
						WillReturnRows(rowsInstance)
				}
				{
					var (
						arg1 = 1
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `users`.`id` FROM `users` WHERE `users`.`role_id` = ?",
						),
					).
						WithArgs(arg1).
						WillReturnError(errors.New("error something"))
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_delete_error",
			input: args{
				db: mockDB,
			},
			output: result{
				err: nil,
			},
			before: func() {
				{
					var (
						arg1         = "super_admin"
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `roles`.`id` FROM `roles` WHERE `roles`.`slug` = ? ORDER BY `roles`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1).
						WillReturnRows(rowsInstance)
				}
				{
					var (
						arg1 = 1
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `users`.`id` FROM `users` WHERE `users`.`role_id` = ?",
						),
					).
						WithArgs(arg1).
						WillReturnRows(rowsInstance)
				}
				{
					var (
						arg1 = 1
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							"DELETE FROM `users` WHERE `users`.`id` = ?",
						),
					).
						WithArgs(arg1).
						WillReturnError(errors.New("error delete"))
					mockQuery.ExpectRollback()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_clear_user",
			input: args{
				db: mockDB,
			},
			output: result{
				err: nil,
			},
			before: func() {
				{
					var (
						arg1         = "super_admin"
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `roles`.`id` FROM `roles` WHERE `roles`.`slug` = ? ORDER BY `roles`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1).
						WillReturnRows(rowsInstance)
				}
				{
					var (
						arg1 = 1
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `users`.`id` FROM `users` WHERE `users`.`role_id` = ?",
						),
					).
						WithArgs(arg1).
						WillReturnRows(rowsInstance)
				}
				{
					var (
						arg1 = 1
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							"DELETE FROM `users` WHERE `users`.`id` = ?",
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

			err := userDatabaseSeederSQLRepository.Clear(tC.input.db)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
