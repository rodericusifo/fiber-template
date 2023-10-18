package resource

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func init() {
	SetupTestPermissionResource()
}

func TestPermissionResource_FirstPermission(t *testing.T) {
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
						result *sql.Permission = nil
						err    error     = errors.New("error something")
					)
					mockPermissionDatabaseSQLRepository.EXPECT().FirstPermission(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_first_permission",
			input: args{
				query: nil,
			},
			output: result{
				value: &sql.Permission{
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
						result *sql.Permission = &sql.Permission{
							ID:        3,
							XID:       mockUUID,
							Name:      "Super Admin",
							Slug:      "super_admin",
							CreatedAt: mockDateTime,
							UpdatedAt: mockDateTime,
						}
						err error = nil
					)
					mockPermissionDatabaseSQLRepository.EXPECT().FirstPermission(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, err := permissionResource.FirstPermission(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
