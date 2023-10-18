package user

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
)

func init() {
	SetupTestMysqlUserDatabaseSQLRepository()
}

func TestMysqlUserDatabaseSQLRepository_SaveUser(t *testing.T) {
	type (
		args struct {
			payload *sql.User
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
			desc: "[ERROR]_because_missing_email",
			input: args{
				payload: &sql.User{
					Name:     "Someone",
					Password: mockHashPassword,
					RoleID:   1,
				},
			},
			output: result{
				err: errors.New("missing email user"),
			},
			before: func() {
				{
					monkey.Patch(uuid.NewString, func() string {
						return mockUUID
					})
				}
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDateTime
					})
				}
				{
					var (
						arg1 = mockUUID
						arg2 = "Someone"
						arg3 = ""
						arg4 = mockHashPassword
						arg5 = 1
						arg6 = mockDateTime
						arg7 = mockDateTime
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							"INSERT INTO `users` (`xid`,`name`,`email`,`password`,`role_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)",
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7).
						WillReturnError(errors.New("missing email user"))
					mockQuery.ExpectRollback()
				}
			},
			after: func() {
				{
					monkey.Unpatch(uuid.NewString)
				}
				{
					monkey.Unpatch(time.Now)
				}
			},
		},
		{
			desc: "[SUCCESS]_success_save_user",
			input: args{
				payload: &sql.User{
					Name:     "Someone",
					Email:    "someone@mail.com",
					Password: mockHashPassword,
					RoleID:   1,
				},
			},
			output: result{
				err: nil,
			},
			before: func() {
				{
					monkey.Patch(uuid.NewString, func() string {
						return mockUUID
					})
				}
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDateTime
					})
				}
				{
					var (
						arg1 = mockUUID
						arg2 = "Someone"
						arg3 = "someone@mail.com"
						arg4 = mockHashPassword
						arg5 = 1
						arg6 = mockDateTime
						arg7 = mockDateTime
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							"INSERT INTO `users` (`xid`,`name`,`email`,`password`,`role_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)",
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7).
						WillReturnResult(sqlmock.NewResult(0, 1))
					mockQuery.ExpectCommit()
				}
			},
			after: func() {
				{
					monkey.Unpatch(uuid.NewString)
				}
				{
					monkey.Unpatch(time.Now)
				}
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			err := userDatabaseSQLRepository.SaveUser(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
