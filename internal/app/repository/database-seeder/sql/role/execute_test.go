package role

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

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
)

func SetupTestExecuteMysqlRoleDatabaseSeederSQLRepository() {
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
	SetupTestExecuteMysqlRoleDatabaseSeederSQLRepository()
}

func TestExecuteMysqlRoleDatabaseSeederRepository(t *testing.T) {
	type (
		args struct {
			isRebuildData config.IsRebuildDataDBSeederMysqlRole
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
			desc: "[ERROR]_error_clear_roles",
			input: args{
				db:            mockDB,
				isRebuildData: true,
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
			desc: "[ERROR]_success_clear_roles_and_fail_seed_role",
			input: args{
				db:            mockDB,
				isRebuildData: true,
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
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDateTime
					})
				}
				{
					var (
						arg1 = "77ce5f5f-2db6-4fff-b6e2-87464e0a9608"
						arg2 = "super_admin"
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `roles`.`id` FROM `roles` WHERE `roles`.`xid` = ? AND `roles`.`slug` = ? ORDER BY `roles`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(gorm.ErrRecordNotFound)
				}
				{
					var (
						arg1 = "77ce5f5f-2db6-4fff-b6e2-87464e0a9608"
						arg2 = "Super Admin"
						arg3 = "super_admin"
						arg4 = mockDateTime
						arg5 = mockDateTime
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							"INSERT INTO `roles` (`xid`,`name`,`slug`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)",
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5).
						WillReturnError(errors.New("error create batches"))
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
			desc: "[SUCCESS]_success_clear_roles_and_seed_roles",
			input: args{
				db:            mockDB,
				isRebuildData: true,
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
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDateTime
					})
				}
				{
					var (
						arg1 = "77ce5f5f-2db6-4fff-b6e2-87464e0a9608"
						arg2 = "super_admin"
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `roles`.`id` FROM `roles` WHERE `roles`.`xid` = ? AND `roles`.`slug` = ? ORDER BY `roles`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(gorm.ErrRecordNotFound)
				}
				{
					var (
						arg1 = "77ce5f5f-2db6-4fff-b6e2-87464e0a9608"
						arg2 = "Super Admin"
						arg3 = "super_admin"
						arg4 = mockDateTime
						arg5 = mockDateTime
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							"INSERT INTO `roles` (`xid`,`name`,`slug`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)",
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5).
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

			ExecuteMysqlRoleDatabaseSeederRepository(tC.input.isRebuildData, tC.input.db)

			tC.after()
		})
	}
}
