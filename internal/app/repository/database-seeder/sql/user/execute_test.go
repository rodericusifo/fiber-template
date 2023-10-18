package user

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/pkg/config"
	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/mocker"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/patcher"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
)

func SetupTestExecuteMysqlUserDatabaseSeederSQLRepository() {
	dialect := pkg_constant.MYSQL
	db, mock := mocker.MockDatabaseSQLConnection(dialect)

	mockQuery = mock
	mockDB = db

	layoutFormat := constant.DEFAULT_TIME_LAYOUT.(string)

	mockDateString = "2015-09-02 08:04:00"
	mockDateTime, _ = time.Parse(layoutFormat, mockDateString)

	mockHashPassword = "$2y$14$rnbG3JhbftD.iQV0QRf5GeNI/XlI85KF2kzrf4hnOs48cSoqPvsmG"
}

func init() {
	SetupTestExecuteMysqlUserDatabaseSeederSQLRepository()
}

func TestExecuteMysqlUserDatabaseSeederRepository(t *testing.T) {
	type (
		args struct {
			isRebuildData config.IsRebuildDataDBSeederMysqlUser
			db            config.MysqlDatabaseSQLConnection
		}
	)

	testCases := []struct {
		desc   string
		input  args
		before func()
		after  func()
	}{
		{
			desc: "[ERROR]_error_clear_users",
			input: args{
				db:            mockDB,
				isRebuildData: true,
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
			desc: "[ERROR]_success_clear_users_and_fail_seed_user",
			input: args{
				db:            mockDB,
				isRebuildData: true,
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
						arg1         = 1
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
			after: func() {
				{
					monkey.Unpatch(time.Now)
				}
			},
		},
		{
			desc: "[SUCCESS]_success_clear_users_and_seed_users",
			input: args{
				db:            mockDB,
				isRebuildData: true,
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
						arg1         = 1
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

			ExecuteMysqlUserDatabaseSeederRepository(tC.input.isRebuildData, tC.input.db)

			tC.after()
		})
	}
}
