package service

import (
	"errors"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/patcher"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func init() {
	SetupTestAuthService()
}

func TestAuthService_RegisterAuth(t *testing.T) {
	type (
		args struct {
			payload *input.RegisterAuthDTO
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
			desc: "[ERROR]_because_error_something_when_get_user",
			input: args{
				payload: &input.RegisterAuthDTO{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: "abc123",
					Role:     constant.NON_ADMIN,
				},
			},
			output: result{
				err: errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Selects: []pkg_types.SelectQuerySQLOperation{
								{Field: "id"},
							},
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "email", Operator: "=", Value: "Ifo@gmail.com"},
								},
							},
						}
					)
					var (
						result *sql.User = nil
						err    error     = errors.New("error something")
					)
					mockUserResource.On("FirstUser", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_user_already_registered",
			input: args{
				payload: &input.RegisterAuthDTO{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: "abc123",
					Role:     constant.NON_ADMIN,
				},
			},
			output: result{
				err: fiber.NewError(fiber.StatusConflict, "user already registered"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Selects: []pkg_types.SelectQuerySQLOperation{
								{Field: "id"},
							},
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "email", Operator: "=", Value: "Ifo@gmail.com"},
								},
							},
						}
					)
					var (
						result *sql.User = &sql.User{
							ID: 1,
						}
						err error = nil
					)
					mockUserResource.On("FirstUser", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_error_something_when_hash_user_password",
			input: args{
				payload: &input.RegisterAuthDTO{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: "abc123",
					Role:     constant.NON_ADMIN,
				},
			},
			output: result{
				err: errors.New("error something"),
			},
			before: func() {
				{
					patcher.GenerateHashFromPassword = func(password string) (string, error) {
						return "", errors.New("error something")
					}
				}
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Selects: []pkg_types.SelectQuerySQLOperation{
								{Field: "id"},
							},
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "email", Operator: "=", Value: "Ifo@gmail.com"},
								},
							},
						}
					)
					var (
						result *sql.User = nil
						err    error     = gorm.ErrRecordNotFound
					)
					mockUserResource.On("FirstUser", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_error_something_when_create_user",
			input: args{
				payload: &input.RegisterAuthDTO{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: "abc123",
					Role:     constant.NON_ADMIN,
				},
			},
			output: result{
				err: errors.New("error something"),
			},
			before: func() {
				{
					patcher.GenerateHashFromPassword = func(password string) (string, error) {
						return mockHashPassword, nil
					}
				}
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Selects: []pkg_types.SelectQuerySQLOperation{
								{Field: "id"},
							},
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "email", Operator: "=", Value: "Ifo@gmail.com"},
								},
							},
						}
					)
					var (
						result *sql.User = nil
						err    error     = gorm.ErrRecordNotFound
					)
					mockUserResource.On("FirstUser", arg1).Return(result, err).Once()
				}
				{
					var (
						arg1 *sql.User = &sql.User{
							Name:     "Ifo",
							Email:    "Ifo@gmail.com",
							Password: mockHashPassword,
							Role:     constant.NON_ADMIN,
						}
					)
					var (
						err error = errors.New("error something")
					)
					mockUserResource.On("SaveUser", arg1).Return(err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_register_auth",
			input: args{
				payload: &input.RegisterAuthDTO{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: "abc123",
					Role:     constant.NON_ADMIN,
				},
			},
			output: result{
				err: nil,
			},
			before: func() {
				{
					patcher.GenerateHashFromPassword = func(password string) (string, error) {
						return mockHashPassword, nil
					}
				}
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Selects: []pkg_types.SelectQuerySQLOperation{
								{Field: "id"},
							},
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "email", Operator: "=", Value: "Ifo@gmail.com"},
								},
							},
						}
					)
					var (
						result *sql.User = nil
						err    error     = gorm.ErrRecordNotFound
					)
					mockUserResource.On("FirstUser", arg1).Return(result, err).Once()
				}
				{
					var (
						arg1 *sql.User = &sql.User{
							Name:     "Ifo",
							Email:    "Ifo@gmail.com",
							Password: mockHashPassword,
							Role:     constant.NON_ADMIN,
						}
					)
					var (
						err error = nil
					)
					mockUserResource.On("SaveUser", arg1).Return(err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			err := authService.RegisterAuth(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
