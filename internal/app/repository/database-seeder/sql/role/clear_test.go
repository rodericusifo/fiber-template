package role

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func init() {
	SetupTestRoleDatabaseSeederSQLRepository()
}

func TestRoleDatabaseSeederSQLRepository_Clear(t *testing.T) {
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
				RoleSeedData = []*RoleSeedPayload{
					{
						XID:  "77ce5f5f-2db6-4fff-b6e2-87464e0a9608",
						Name: "Super Admin",
						Slug: "super_admin",
					},
				}
				{
					var (
						arg1 ="super_admin"
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `roles`.`id` FROM `roles` WHERE `roles`.`slug` IN (?)",
						),
					).
						WithArgs(arg1).
						WillReturnError(errors.New("error something"))
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR_IN_LOOP]_because_error_when_delete_roles",
			input: args{
				db: mockDB,
			},
			output: result{
				err: nil,
			},
			before: func() {
				RoleSeedData = []*RoleSeedPayload{
					{
						XID:  "77ce5f5f-2db6-4fff-b6e2-87464e0a9608",
						Name: "Super Admin",
						Slug: "super_admin",
					},
				}
				{
					var (
						arg1 ="super_admin"
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `roles`.`id` FROM `roles` WHERE `roles`.`slug` IN (?)",
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
							"DELETE FROM `roles` WHERE `roles`.`id` = ?",
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
			desc: "[SUCCESS]_clear_roles",
			input: args{
				db: mockDB,
			},
			output: result{
				err: nil,
			},
			before: func() {
				RoleSeedData = []*RoleSeedPayload{
					{
						XID:  "77ce5f5f-2db6-4fff-b6e2-87464e0a9608",
						Name: "Super Admin",
						Slug: "super_admin",
					},
				}
				{
					var (
						arg1 ="super_admin"
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `roles`.`id` FROM `roles` WHERE `roles`.`slug` IN (?)",
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
							"DELETE FROM `roles` WHERE `roles`.`id` = ?",
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

			err := roleDatabaseSeederSQLRepository.Clear(tC.input.db)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
