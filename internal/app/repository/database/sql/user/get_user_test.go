package user

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/pkg/constant"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func init() {
	SetupTestMysqlUserDatabaseSQLRepository()
}

func TestMysqlUserDatabaseSQLRepository_GetUser(t *testing.T) {
	type (
		args struct {
			query *pkg_types.QuerySQL
		}
		result struct {
			value *sql.User
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
					Selects: []pkg_types.SelectQuerySQLOperation{
						{Field: "id"},
						{Field: "xid"},
						{Field: "name"},
						{Field: "email"},
						{Field: "password"},
						{Field: "role"},
						{Field: "created_at"},
						{Field: "updated_at"},
					},
					Searches: [][]pkg_types.SearchQuerySQLOperation{
						{
							{Field: "email", Operator: "=", Value: "someone@mail.com"},
						},
					},
				},
			},
			output: result{
				value: nil,
				err:   errors.New("something error"),
			},
			before: func() {
				{
					var (
						arg1 = "someone@mail.com"
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `users`.`id`,`users`.`xid`,`users`.`name`,`users`.`email`,`users`.`password`,`users`.`role`,`users`.`created_at`,`users`.`updated_at` FROM `users` WHERE `users`.`email` = ? ORDER BY `users`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1).
						WillReturnError(errors.New("something error"))
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_get_user",
			input: args{
				query: &pkg_types.QuerySQL{
					Selects: []pkg_types.SelectQuerySQLOperation{
						{Field: "id"},
						{Field: "xid"},
						{Field: "name"},
						{Field: "email"},
						{Field: "password"},
						{Field: "role"},
						{Field: "created_at"},
						{Field: "updated_at"},
					},
					Searches: [][]pkg_types.SearchQuerySQLOperation{
						{
							{Field: "email", Operator: "=", Value: "someone@mail.com"},
						},
					},
				},
			},
			output: result{
				value: &sql.User{
					ID:        3,
					XID:       mockUUID,
					Name:      "Someone",
					Email:     "someone@mail.com",
					Password:  mockHashPassword,
					Role:      constant.ADMIN,
					CreatedAt: mockDateTime,
					UpdatedAt: mockDateTime,
				},
				err: nil,
			},
			before: func() {
				{
					var (
						arg1         = "someone@mail.com"
						rowsInstance = sqlmock.NewRows([]string{"id", "xid", "name", "email", "password", "role", "created_at", "updated_at"})
					)
					rowsInstance.AddRow(3, mockUUID, "Someone", "someone@mail.com", mockHashPassword, constant.ADMIN, mockDateTime, mockDateTime)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `users`.`id`,`users`.`xid`,`users`.`name`,`users`.`email`,`users`.`password`,`users`.`role`,`users`.`created_at`,`users`.`updated_at` FROM `users` WHERE `users`.`email` = ? ORDER BY `users`.`id` LIMIT 1",
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

			value, err := userDatabaseSQLRepository.GetUser(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
