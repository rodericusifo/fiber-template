package user

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/pkg/util/patcher"
)

func init() {
	SetupTestUserDatabaseSeederSQLRepository()
}

func TestUserDatabaseSeederSQLRepository_Seed(t *testing.T) {
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
			desc: "[ERROR]_because_role_not_found",
			input: args{
				db: mockDB,
			},
			output: result{
				err: errors.New("role not found"),
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
						WillReturnError(errors.New("role not found"))
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR_IN_LOOP]_because_validation",
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
				UserSeedData = []*UserSeedPayload{
					{
						Name:     "admin",
						Email:    "admin@gmail.com",
						Password: "p4ssw0rd",
					},
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR_IN_LOOP]_because_get_user_error",
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
				UserSeedData = []*UserSeedPayload{
					{
						XID:      "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1",
						Name:     "admin",
						Email:    "admin@gmail.com",
						Password: "p4ssw0rd",
					},
				}
				{
					var (
						arg1 = "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1"
						arg2 = "admin@gmail.com"
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `users`.`id` FROM `users` WHERE `users`.`xid` = ? AND `users`.`email` = ? ORDER BY `users`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(errors.New("unexpected error"))
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR_IN_LOOP]_because_user_already_registered",
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
				UserSeedData = []*UserSeedPayload{
					{
						XID:      "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1",
						Name:     "admin",
						Email:    "admin@gmail.com",
						Password: "p4ssw0rd",
					},
				}
				{
					var (
						arg1         = "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1"
						arg2         = "admin@gmail.com"
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `users`.`id` FROM `users` WHERE `users`.`xid` = ? AND `users`.`email` = ? ORDER BY `users`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnRows(rowsInstance)
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR_IN_LOOP]_because_error_hash_password",
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
				UserSeedData = []*UserSeedPayload{
					{
						XID:      "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1",
						Name:     "admin",
						Email:    "admin@gmail.com",
						Password: "p4ssw0rd",
					},
				}
				{
					patcher.GenerateHashFromPassword = func(password string) (string, error) {
						return "", errors.New("fail hash password")
					}
				}
				{
					var (
						arg1 = "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1"
						arg2 = "admin@gmail.com"
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `users`.`id` FROM `users` WHERE `users`.`xid` = ? AND `users`.`email` = ? ORDER BY `users`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(gorm.ErrRecordNotFound)
				}
			},
			after: func() {},
		},
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
				UserSeedData = []*UserSeedPayload{
					{
						XID:      "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1",
						Name:     "admin",
						Email:    "admin@gmail.com",
						Password: "p4ssw0rd",
					},
				}
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDateTime
					})
				}
				{
					patcher.GenerateHashFromPassword = func(password string) (string, error) {
						return mockHashPassword, nil
					}
				}
				{
					var (
						arg1 = "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1"
						arg2 = "admin@gmail.com"
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `users`.`id` FROM `users` WHERE `users`.`xid` = ? AND `users`.`email` = ? ORDER BY `users`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(gorm.ErrRecordNotFound)
				}
				{
					var (
						arg1 = "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1"
						arg2 = "admin"
						arg3 = "admin@gmail.com"
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
						WillReturnError(errors.New("error something"))
					mockQuery.ExpectRollback()
				}
			},
			after: func() {
				{
					monkey.Unpatch(time.Now)
				}
			},
		},
		{
			desc: "[SUCCESS]_success_seed_users",
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
				UserSeedData = []*UserSeedPayload{
					{
						XID:      "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1",
						Name:     "admin",
						Email:    "admin@gmail.com",
						Password: "p4ssw0rd",
					},
				}
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDateTime
					})
				}
				{
					patcher.GenerateHashFromPassword = func(password string) (string, error) {
						return mockHashPassword, nil
					}
				}
				{
					var (
						arg1 = "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1"
						arg2 = "admin@gmail.com"
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `users`.`id` FROM `users` WHERE `users`.`xid` = ? AND `users`.`email` = ? ORDER BY `users`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(gorm.ErrRecordNotFound)
				}
				{
					var (
						arg1 = "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1"
						arg2 = "admin"
						arg3 = "admin@gmail.com"
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
					monkey.Unpatch(time.Now)
				}
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			err := userDatabaseSeederSQLRepository.Seed(tC.input.db)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
