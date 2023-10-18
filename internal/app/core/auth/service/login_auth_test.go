package service

import (
	"errors"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/app/core/auth/service/dto/output"
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/pkg/types"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/patcher"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func init() {
	SetupTestAuthService()
}

func TestAuthService_LoginAuth(t *testing.T) {
	type (
		args struct {
			payload *input.LoginAuthDTO
		}
		result struct {
			value *output.LoginAuthDTO
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
			desc: "[ERROR]_because_user_not_found",
			input: args{
				payload: &input.LoginAuthDTO{
					Email:    "someone@mail.com",
					Password: mockPassword,
				},
			},
			output: result{
				value: nil,
				err:   fiber.NewError(fiber.StatusNotFound, "user not found"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
								},
							},
						}
					)
					var (
						result *sql.User = nil
						err    error     = gorm.ErrRecordNotFound
					)
					mockUserResource.EXPECT().FirstUser(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_something_error_happens",
			input: args{
				payload: &input.LoginAuthDTO{
					Email:    "someone@mail.com",
					Password: mockPassword,
				},
			},
			output: result{
				value: nil,
				err:   errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
								},
							},
						}
					)
					var (
						result *sql.User = nil
						err    error     = errors.New("error something")
					)
					mockUserResource.EXPECT().FirstUser(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_email_and_password_not_match",
			input: args{
				payload: &input.LoginAuthDTO{
					Email:    "someone@mail.com",
					Password: "123",
				},
			},
			output: result{
				value: nil,
				err:   fiber.NewError(fiber.StatusUnauthorized, "email and password not match"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
								},
							},
						}
					)
					var (
						result *sql.User = &sql.User{
							ID:        3,
							XID:       mockUUID,
							Name:      "Someone",
							Email:     "someone@mail.com",
							Password:  mockHashPassword,
							RoleID:    1,
							CreatedAt: mockDateTime,
							UpdatedAt: mockDateTime,
						}
						err error = nil
					)
					mockUserResource.EXPECT().FirstUser(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_generate_token_from_claims_failed",
			input: args{
				payload: &input.LoginAuthDTO{
					Email:    "someone@mail.com",
					Password: mockPassword,
				},
			},
			output: result{
				value: nil,
				err:   errors.New("error generate token from claims"),
			},
			before: func() {
				{
					patcher.GenerateJWTTokenFromClaims = func(claims *types.JwtCustomClaims) (string, error) {
						return "", errors.New("error generate token from claims")
					}
				}
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
								},
							},
						}
					)
					var (
						result *sql.User = &sql.User{
							ID:        3,
							XID:       mockUUID,
							Name:      "Someone",
							Email:     "someone@mail.com",
							Password:  mockHashPassword,
							RoleID:    1,
							CreatedAt: mockDateTime,
							UpdatedAt: mockDateTime,
						}
						err error = nil
					)
					mockUserResource.EXPECT().FirstUser(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_login_auth",
			input: args{
				payload: &input.LoginAuthDTO{
					Email:    "someone@mail.com",
					Password: mockPassword,
				},
			},
			output: result{
				value: &output.LoginAuthDTO{
					Token: mockJWTToken,
				},
				err: nil,
			},
			before: func() {
				{
					patcher.GenerateJWTTokenFromClaims = func(claims *types.JwtCustomClaims) (string, error) {
						return mockJWTToken, nil
					}
				}
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
								},
							},
						}
					)
					var (
						result *sql.User = &sql.User{
							ID:        3,
							XID:       mockUUID,
							Name:      "Someone",
							Email:     "someone@mail.com",
							Password:  mockHashPassword,
							RoleID:    1,
							CreatedAt: mockDateTime,
							UpdatedAt: mockDateTime,
						}
						err error = nil
					)
					mockUserResource.EXPECT().FirstUser(arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, err := authService.LoginAuth(tC.input.payload)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
