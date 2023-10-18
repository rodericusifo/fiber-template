package role

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func init() {
	SetupTestRoleDatabaseSeederSQLRepository()
}

func TestRoleDatabaseSeederSQLRepository_Seed(t *testing.T) {
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
			desc: "[ERROR_IN_LOOP]_because_validation",
			input: args{
				db: mockDB,
			},
			output: result{
				err: nil,
			},
			before: func() {
				RoleSeedData = []*RoleSeedPayload{
					{
						Name: "Super Admin",
					},
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR_IN_LOOP]_because_error_someting_when_get_role",
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
					monkey.Patch(time.Now, func() time.Time {
						return mockDateTime
					})
				}
				{
					var (
						arg1 = "77ce5f5f-2db6-4fff-b6e2-87464e0a9608"
						arg2 = "super_admin"
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `roles`.`id` FROM `roles` WHERE `roles`.`xid` = ? AND `roles`.`slug` = ? ORDER BY `roles`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(errors.New("error something"))
				}
			},
			after: func() {
				{
					monkey.Unpatch(time.Now)
				}
			},
		},
		{
			desc: "[ERROR_IN_LOOP]_because_role_already_registered",
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
						WillReturnRows(rowsInstance)
				}
			},
			after: func() {
				{
					monkey.Unpatch(time.Now)
				}
			},
		},
		{
			desc: "[ERROR]_error_when_seed_batches_role",
			input: args{
				db: mockDB,
			},
			output: result{
				err: errors.New("error create batches"),
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
			desc: "[SUCCESS]_seed_role",
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

			err := roleDatabaseSeederSQLRepository.Seed(tC.input.db)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
