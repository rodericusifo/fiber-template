package employee

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
	SetupTestMysqlEmployeeDatabaseSQLRepository()
}

func TestMysqlEmployeeDatabaseSQLRepository_SaveEmployee(t *testing.T) {
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
			desc: "[ERROR]_because_missing_email",
			input: args{
				payload: &sql.Employee{
					Name:   "Someone",
					UserID: 1,
				},
			},
			output: result{
				err: errors.New("missing email employee"),
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
						arg1      = mockUUID
						arg2      = "Someone"
						arg3      = ""
						arg4  any = nil
						arg5  any = nil
						arg6  any = nil
						arg7      = 1
						arg8      = mockDateTime
						arg9      = mockDateTime
						arg10 any = nil
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							"INSERT INTO `employees` (`xid`,`name`,`email`,`address`,`age`,`birthday`,`user_id`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?,?,?)",
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10).
						WillReturnError(errors.New("missing email employee"))
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
			desc: "[SUCCESS]_success_create_employee",
			input: args{
				payload: &sql.Employee{
					Name:   "Someone",
					Email:  "someone@mail.com",
					UserID: 1,
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
						arg1      = mockUUID
						arg2      = "Someone"
						arg3      = "someone@mail.com"
						arg4  any = nil
						arg5  any = nil
						arg6  any = nil
						arg7      = 1
						arg8      = mockDateTime
						arg9      = mockDateTime
						arg10 any = nil
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							"INSERT INTO `employees` (`xid`,`name`,`email`,`address`,`age`,`birthday`,`user_id`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?,?,?)",
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10).
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
		{
			desc: "[ERROR]_because_something_error_happens",
			input: args{
				payload: &sql.Employee{
					ID:        1,
					XID:       mockUUID,
					Name:      "Someone",
					Email:     "someone@mail.com",
					Address:   &mockAddress,
					Age:       &mockAge,
					Birthday:  &mockBirthdayTime,
					UserID:    1,
					CreatedAt: mockDateTime,
					UpdatedAt: mockDateTime,
				},
			},
			output: result{
				err: errors.New("error something"),
			},
			before: func() {
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDateTime
					})
				}
				{
					var (
						arg1      = mockUUID
						arg2      = "Someone"
						arg3      = "someone@mail.com"
						arg4      = &mockAddress
						arg5      = &mockAge
						arg6      = &mockBirthdayTime
						arg7      = 1
						arg8      = mockDateTime
						arg9      = mockDateTime
						arg10 any = nil
						arg11     = 1
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							"UPDATE `employees` SET `xid`=?,`name`=?,`email`=?,`address`=?,`age`=?,`birthday`=?,`user_id`=?,`created_at`=?,`updated_at`=?,`deleted_at`=? WHERE `employees`.`deleted_at` IS NULL AND `id` = ?",
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11).
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
			desc: "[SUCCESS]_success_update_employee",
			input: args{
				payload: &sql.Employee{
					ID:        1,
					XID:       mockUUID,
					Name:      "Someone",
					Email:     "someone@mail.com",
					Address:   &mockAddress,
					Age:       &mockAge,
					Birthday:  &mockBirthdayTime,
					UserID:    1,
					CreatedAt: mockDateTime,
					UpdatedAt: mockDateTime,
				},
			},
			output: result{
				err: nil,
			},
			before: func() {
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDateTime
					})
				}
				{
					var (
						arg1      = mockUUID
						arg2      = "Someone"
						arg3      = "someone@mail.com"
						arg4      = &mockAddress
						arg5      = &mockAge
						arg6      = &mockBirthdayTime
						arg7      = 1
						arg8      = mockDateTime
						arg9      = mockDateTime
						arg10 any = nil
						arg11     = 1
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							"UPDATE `employees` SET `xid`=?,`name`=?,`email`=?,`address`=?,`age`=?,`birthday`=?,`user_id`=?,`created_at`=?,`updated_at`=?,`deleted_at`=? WHERE `employees`.`deleted_at` IS NULL AND `id` = ?",
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11).
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

			err := employeeDatabaseSQLRepository.SaveEmployee(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
