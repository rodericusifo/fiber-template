package permission

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func init() {
	SetupTestMysqlPermissionDatabaseSQLRepository()
}

func TestMysqlPermissionDatabaseSQLRepository_FirstPermission(t *testing.T) {
	type (
		args struct {
			query *pkg_types.QuerySQL
		}
		result struct {
			value *sql.Permission
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
						{Field: "slug"},
						{Field: "created_at"},
						{Field: "updated_at"},
					},
					Searches: [][]pkg_types.SearchQuerySQLOperation{
						{
							{Field: "xid", Operator: "=", Value: mockUUID},
							{Field: "slug", Operator: "=", Value: "super_admin"},
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
						arg1 = mockUUID
						arg2 = "super_admin"
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `permissions`.`id`,`permissions`.`xid`,`permissions`.`name`,`permissions`.`slug`,`permissions`.`created_at`,`permissions`.`updated_at` FROM `permissions` WHERE `permissions`.`xid` = ? AND `permissions`.`slug` = ? ORDER BY `permissions`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(errors.New("something error"))
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_first_permission",
			input: args{
				query: &pkg_types.QuerySQL{
					Selects: []pkg_types.SelectQuerySQLOperation{
						{Field: "id"},
						{Field: "xid"},
						{Field: "name"},
						{Field: "slug"},
						{Field: "created_at"},
						{Field: "updated_at"},
					},
					Searches: [][]pkg_types.SearchQuerySQLOperation{
						{
							{Field: "xid", Operator: "=", Value: mockUUID},
							{Field: "slug", Operator: "=", Value: "super_admin"},
						},
					},
				},
			},
			output: result{
				value: &sql.Permission{
					ID: 1,
					XID: mockUUID,
					Name: "Super Admin",
					Slug: "super_admin",
					CreatedAt: mockDateTime,
					UpdatedAt: mockDateTime,
				},
				err:   nil,
			},
			before: func() {
				{
					var (
						arg1 = mockUUID
						arg2 = "super_admin"
						rowsInstance = sqlmock.NewRows([]string{"id", "xid", "name", "slug", "created_at", "updated_at"})
					)
					rowsInstance.AddRow(1, mockUUID, "Super Admin", "super_admin",mockDateTime, mockDateTime)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `permissions`.`id`,`permissions`.`xid`,`permissions`.`name`,`permissions`.`slug`,`permissions`.`created_at`,`permissions`.`updated_at` FROM `permissions` WHERE `permissions`.`xid` = ? AND `permissions`.`slug` = ? ORDER BY `permissions`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnRows(rowsInstance)
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, err := permissionDatabaseSQLRepository.FirstPermission(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
