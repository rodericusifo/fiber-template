package employee

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func init() {
	SetupTestMysqlEmployeeDatabaseSQLRepository()
}

func TestMysqlEmployeeDatabaseSQLRepository_FindEmployees(t *testing.T) {
	type (
		args struct {
			query *pkg_types.QuerySQL
		}
		result struct {
			value []*sql.Employee
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
						{Field: "address"},
						{Field: "age"},
						{Field: "birthday"},
						{Field: "created_at"},
						{Field: "updated_at"},
					},
					Searches: [][]pkg_types.SearchQuerySQLOperation{
						{
							{Field: "name", Operator: "LIKE", Value: fmt.Sprint("%", "someone", "%")},
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
					Joins: []pkg_types.JoinQuerySQLOperation{
						{
							Relation: "User",
						},
					},
					Orders: []pkg_types.OrderQuerySQLOperation{
						{Field: "name"},
						{Field: "age", Descending: true},
					},
					Groups: []pkg_types.GroupQuerySQLOperation{
						{Field: "name"},
					},
					Limit:       10,
					Offset:      10,
					WithDeleted: true,
				},
			},
			output: result{
				value: nil,
				err:   errors.New("something error"),
			},
			before: func() {
				{
					var (
						arg1 = fmt.Sprint("%", "someone", "%")
						arg2 = 1
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `employees`.`id`,`employees`.`xid`,`employees`.`name`,`employees`.`email`,`employees`.`address`,`employees`.`age`,`employees`.`birthday`,`employees`.`created_at`,`employees`.`updated_at`,`User`.`id` AS `User__id`,`User`.`xid` AS `User__xid`,`User`.`name` AS `User__name`,`User`.`email` AS `User__email`,`User`.`password` AS `User__password`,`User`.`role_id` AS `User__role_id`,`User`.`created_at` AS `User__created_at`,`User`.`updated_at` AS `User__updated_at` FROM `employees` LEFT JOIN `users` `User` ON `employees`.`user_id` = `User`.`id` WHERE `employees`.`name` LIKE ? AND `employees`.`user_id` = ? GROUP BY `employees`.`name` ORDER BY `employees`.`name`,`employees`.`age` DESC LIMIT 10 OFFSET 10",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(errors.New("something error"))
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_something_error_happens_1",
			input: args{
				query: &pkg_types.QuerySQL{
					Selects: []pkg_types.SelectQuerySQLOperation{
						{Field: "id"},
						{Field: "xid"},
						{Field: "name"},
						{Field: "email"},
						{Field: "address"},
						{Field: "age"},
						{Field: "birthday"},
						{Field: "created_at"},
						{Field: "updated_at"},
					},
					Searches: [][]pkg_types.SearchQuerySQLOperation{
						{
							{Field: "name", Operator: "LIKE", Value: fmt.Sprint("%", "someone", "%")},
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
					Joins: []pkg_types.JoinQuerySQLOperation{
						{
							Relation: "User",
							Selects: []pkg_types.SelectJoinQuerySQLOperation{
								{Field: "id"},
								{Field: "xid"},
								{Field: "name"},
							},
						},
					},
					Limit:  10,
					Offset: 10,
				},
			},
			output: result{
				value: nil,
				err:   errors.New("something error"),
			},
			before: func() {
				{
					var (
						arg1 = fmt.Sprint("%", "someone", "%")
						arg2 = 1
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `employees`.`id`,`employees`.`xid`,`employees`.`name`,`employees`.`email`,`employees`.`address`,`employees`.`age`,`employees`.`birthday`,`employees`.`created_at`,`employees`.`updated_at`,`User`.`id` AS `User__id`,`User`.`xid` AS `User__xid`,`User`.`name` AS `User__name` FROM `employees` LEFT JOIN `users` `User` ON `employees`.`user_id` = `User`.`id` WHERE (`employees`.`name` LIKE ? AND `employees`.`user_id` = ?) AND `employees`.`deleted_at` IS NULL LIMIT 10 OFFSET 10",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(errors.New("something error"))
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_something_error_happens_2",
			input: args{
				query: &pkg_types.QuerySQL{
					Selects: []pkg_types.SelectQuerySQLOperation{
						{Field: "id"},
						{Field: "xid"},
						{Field: "name"},
						{Field: "email"},
						{Field: "address"},
						{Field: "age"},
						{Field: "birthday"},
						{Field: "created_at"},
						{Field: "updated_at"},
					},
					Searches: [][]pkg_types.SearchQuerySQLOperation{
						{
							{Field: "name", Operator: "LIKE", Value: fmt.Sprint("%", "someone", "%")},
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
					InnerJoins: []pkg_types.InnerJoinQuerySQLOperation{
						{
							Relation: "User",
							Selects: []pkg_types.SelectJoinQuerySQLOperation{
								{Field: "id"},
								{Field: "xid"},
								{Field: "name"},
							},
						},
					},
					Limit:  10,
					Offset: 10,
				},
			},
			output: result{
				value: nil,
				err:   errors.New("something error"),
			},
			before: func() {
				{
					var (
						arg1 = fmt.Sprint("%", "someone", "%")
						arg2 = 1
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `employees`.`id`,`employees`.`xid`,`employees`.`name`,`employees`.`email`,`employees`.`address`,`employees`.`age`,`employees`.`birthday`,`employees`.`created_at`,`employees`.`updated_at`,`User`.`id` AS `User__id`,`User`.`xid` AS `User__xid`,`User`.`name` AS `User__name` FROM `employees` INNER JOIN `users` `User` ON `employees`.`user_id` = `User`.`id` WHERE (`employees`.`name` LIKE ? AND `employees`.`user_id` = ?) AND `employees`.`deleted_at` IS NULL LIMIT 10 OFFSET 10",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(errors.New("something error"))
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_find_employees",
			input: args{
				query: &pkg_types.QuerySQL{
					Distinct: true,
					Selects: []pkg_types.SelectQuerySQLOperation{
						{Field: "id"},
						{Field: "xid"},
						{Field: "name"},
						{Field: "email"},
						{Field: "address"},
						{Field: "age"},
						{Field: "birthday"},
						{Field: "created_at"},
						{Field: "updated_at"},
					},
					Searches: [][]pkg_types.SearchQuerySQLOperation{
						{
							{Field: "name", Operator: "LIKE", Value: fmt.Sprint("%", "someone", "%")},
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
					Joins: []pkg_types.JoinQuerySQLOperation{
						{
							Relation: "User",
							Selects: []pkg_types.SelectJoinQuerySQLOperation{
								{Field: "id"},
								{Field: "xid"},
								{Field: "name"},
							},
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "name", Operator: "LIKE", Value: fmt.Sprint("%", "sometwo", "%")},
								},
							},
						},
					},
					Limit:  10,
					Offset: 10,
				},
			},
			output: result{
				value: []*sql.Employee{
					{
						ID:        1,
						XID:       mockUUID,
						Name:      "Someone",
						Email:     "someone@mail.com",
						Address:   nil,
						Age:       nil,
						Birthday:  nil,
						CreatedAt: mockDateTime,
						UpdatedAt: mockDateTime,
						User: sql.User{
							ID:   2,
							XID:  mockUUID,
							Name: "sometwo",
						},
					},
				},
				err: nil,
			},
			before: func() {
				{
					var (
						arg1         = fmt.Sprint("%", "sometwo", "%")
						arg2         = fmt.Sprint("%", "someone", "%")
						arg3         = 1
						rowsInstance = sqlmock.NewRows([]string{"id", "xid", "name", "email", "address", "age", "birthday", "created_at", "updated_at", "User__id", "User__xid", "User__name"})
					)
					rowsInstance.AddRow(1, mockUUID, "Someone", "someone@mail.com", nil, nil, nil, mockDateTime, mockDateTime, 2, mockUUID, "sometwo")
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT DISTINCT `employees`.`id`,`employees`.`xid`,`employees`.`name`,`employees`.`email`,`employees`.`address`,`employees`.`age`,`employees`.`birthday`,`employees`.`created_at`,`employees`.`updated_at`,`User`.`id` AS `User__id`,`User`.`xid` AS `User__xid`,`User`.`name` AS `User__name` FROM `employees` LEFT JOIN `users` `User` ON `employees`.`user_id` = `User`.`id` AND `User`.`name` LIKE ? WHERE (`employees`.`name` LIKE ? AND `employees`.`user_id` = ?) AND `employees`.`deleted_at` IS NULL LIMIT 10 OFFSET 10",
						),
					).
						WithArgs(arg1, arg2, arg3).
						WillReturnRows(rowsInstance)
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, err := employeeDatabaseSQLRepository.FindEmployees(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
