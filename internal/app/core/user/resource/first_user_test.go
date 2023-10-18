package resource

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func init() {
	SetupTestUserResource()
}

func TestUserResource_FirstUser(t *testing.T) {
	type (
		args struct {
			query *pkg_types.QuerySQL
		}
		result struct {
			value *sql.User
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
						result *sql.User = nil
						err    error     = errors.New("error something")
					)
					mockUserDatabaseSQLRepository.EXPECT().FirstUser(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_first_user",
			input: args{
				query: nil,
			},
			output: result{
				value: &sql.User{
					ID:        3,
					XID:       mockUUID,
					Name:      "Someone",
					Email:     "someone@mail.com",
					Password:  mockHashPassword,
					RoleID:      1,
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
						result *sql.User = &sql.User{
							ID:        3,
							XID:       mockUUID,
							Name:      "Someone",
							Email:     "someone@mail.com",
							Password:  mockHashPassword,
							RoleID:      1,
							CreatedAt: mockDateTime,
							UpdatedAt: mockDateTime,
						}
						err error = nil
					)
					mockUserDatabaseSQLRepository.EXPECT().FirstUser(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, err := userResource.FirstUser(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
