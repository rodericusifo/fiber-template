package resource

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
)

func init() {
	SetupTestUserResource()
}

func TestUserResource_SaveUser(t *testing.T) {
	type (
		args struct {
			payload *sql.User
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
			desc: "[ERROR]_because_something_error_happens",
			input: args{
				payload: &sql.User{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: mockHashPassword,
					RoleID:   1,
				},
			},
			output: result{
				err: errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 *sql.User = &sql.User{
							Name:     "Ifo",
							Email:    "Ifo@gmail.com",
							Password: mockHashPassword,
							RoleID:   1,
						}
					)
					var (
						err error = errors.New("error something")
					)
					mockUserDatabaseSQLRepository.EXPECT().SaveUser(arg1).Return(err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_save_user",
			input: args{
				payload: &sql.User{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: mockHashPassword,
					RoleID:   1,
				},
			},
			output: result{
				err: nil,
			},
			before: func() {
				{
					var (
						arg1 *sql.User = &sql.User{
							Name:     "Ifo",
							Email:    "Ifo@gmail.com",
							Password: mockHashPassword,
							RoleID:   1,
						}
					)
					var (
						err error = nil
					)
					mockUserDatabaseSQLRepository.EXPECT().SaveUser(arg1).Return(err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			err := userResource.SaveUser(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
