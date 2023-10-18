package resource

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func init() {
	SetupTestRoleResource()
}

func TestRoleResource_FirstRole(t *testing.T) {
	type (
		args struct {
			query *pkg_types.QuerySQL
		}
		result struct {
			value *sql.Role
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
				query: nil,
			},
			output: result{
				value: nil,
				err:   errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = nil
					)
					var (
						result *sql.Role = nil
						err    error     = errors.New("error something")
					)
					mockRoleDatabaseSQLRepository.EXPECT().FirstRole(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_first_role",
			input: args{
				query: nil,
			},
			output: result{
				value: &sql.Role{
					ID:        3,
					XID:       mockUUID,
					Name:      "Super Admin",
					Slug:      "super_admin",
					CreatedAt: mockDateTime,
					UpdatedAt: mockDateTime,
				},
				err: nil,
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = nil
					)
					var (
						result *sql.Role = &sql.Role{
							ID:        3,
							XID:       mockUUID,
							Name:      "Super Admin",
							Slug:      "super_admin",
							CreatedAt: mockDateTime,
							UpdatedAt: mockDateTime,
						}
						err error = nil
					)
					mockRoleDatabaseSQLRepository.EXPECT().FirstRole(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, err := roleResource.FirstRole(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
