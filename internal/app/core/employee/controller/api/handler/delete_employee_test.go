package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/fiber-template/internal/app/core/employee/controller/api/request"
	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
	pkg_util_response "github.com/rodericusifo/fiber-template/pkg/util/response"
)

func init() {
	SetupTestEmployeeHandler()
}

func TestEmployeeHandler_DeleteEmployee(t *testing.T) {
	type (
		args struct {
			requestParams request.DeleteEmployeeRequestParams
		}
		result struct {
			responseStatusCode int
			responseBody       any
		}
	)

	testCases := []struct {
		desc    string
		input   args
		output  result
		before  func()
		after   func()
		isError bool
	}{
		{
			desc: "[ERROR]_because_validation_error",
			input: args{
				requestParams: request.DeleteEmployeeRequestParams{
					XID: mockUUIDV1,
				},
			},
			output: result{
				responseStatusCode: fiber.StatusBadRequest,
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Selects: []pkg_types.SelectQuerySQLOperation{
								{Field: "id"},
								{Field: "name"},
								{Field: "email"},
							},
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "xid", Operator: "=", Value: mockUserXID},
								},
							},
							Joins: []pkg_types.JoinQuerySQLOperation{
								{
									Relation: "Role",
									Selects: []pkg_types.SelectJoinQuerySQLOperation{
										{Field: "id"},
										{Field: "slug"},
									},
								},
							},
						}
					)
					var (
						result *sql.User = &sql.User{
							ID:    1,
							Name:  "super.admin",
							Email: "super.admin@gmail.com",
							Role: sql.Role{
								ID:   1,
								Slug: "super_admin",
							},
						}
						err error = nil
					)
					mockUserResource.EXPECT().FirstUser(arg1).Return(result, err).Once()
				}
			},
			after:   func() {},
			isError: true,
		},
		{
			desc: "[ERROR]_because_unexpected_error_from_service",
			input: args{
				requestParams: request.DeleteEmployeeRequestParams{
					XID: mockUUID,
				},
			},
			output: result{
				responseStatusCode: fiber.StatusInternalServerError,
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Selects: []pkg_types.SelectQuerySQLOperation{
								{Field: "id"},
								{Field: "name"},
								{Field: "email"},
							},
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "xid", Operator: "=", Value: mockUserXID},
								},
							},
							Joins: []pkg_types.JoinQuerySQLOperation{
								{
									Relation: "Role",
									Selects: []pkg_types.SelectJoinQuerySQLOperation{
										{Field: "id"},
										{Field: "slug"},
									},
								},
							},
						}
					)
					var (
						result *sql.User = &sql.User{
							ID:    1,
							Name:  "super.admin",
							Email: "super.admin@gmail.com",
							Role: sql.Role{
								ID:   1,
								Slug: "super_admin",
							},
						}
						err error = nil
					)
					mockUserResource.EXPECT().FirstUser(arg1).Return(result, err).Once()
				}
				{
					var (
						arg1 *input.DeleteEmployeeDTO = &input.DeleteEmployeeDTO{
							XID:    mockUUID,
							UserID: 1,
						}
					)
					var (
						err error = errors.New("unexpected errors")
					)
					mockEmployeeService.EXPECT().DeleteEmployee(arg1).Return(err).Once()
				}
			},
			after:   func() {},
			isError: true,
		},
		{
			desc: "[SUCCESS]_success_delete_employee",
			input: args{
				requestParams: request.DeleteEmployeeRequestParams{
					XID: mockUUID,
				},
			},
			output: result{
				responseStatusCode: fiber.StatusOK,
				responseBody:       pkg_util_response.ResponseSuccess[any]("delete employee success", nil, nil),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.QuerySQL = &pkg_types.QuerySQL{
							Selects: []pkg_types.SelectQuerySQLOperation{
								{Field: "id"},
								{Field: "name"},
								{Field: "email"},
							},
							Searches: [][]pkg_types.SearchQuerySQLOperation{
								{
									{Field: "xid", Operator: "=", Value: mockUserXID},
								},
							},
							Joins: []pkg_types.JoinQuerySQLOperation{
								{
									Relation: "Role",
									Selects: []pkg_types.SelectJoinQuerySQLOperation{
										{Field: "id"},
										{Field: "slug"},
									},
								},
							},
						}
					)
					var (
						result *sql.User = &sql.User{
							ID:    1,
							Name:  "super.admin",
							Email: "super.admin@gmail.com",
							Role: sql.Role{
								ID:   1,
								Slug: "super_admin",
							},
						}
						err error = nil
					)
					mockUserResource.EXPECT().FirstUser(arg1).Return(result, err).Once()
				}
				{
					var (
						arg1 *input.DeleteEmployeeDTO = &input.DeleteEmployeeDTO{
							XID:    mockUUID,
							UserID: 1,
						}
					)
					var (
						err error = nil
					)
					mockEmployeeService.EXPECT().DeleteEmployee(arg1).Return(err).Once()
				}
			},
			after:   func() {},
			isError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer mockApp.Shutdown()

			tC.before()

			url := fmt.Sprintf("/employees/%s/delete", tC.input.requestParams.XID)

			strResponseBodyBytes, _ := json.Marshal(tC.output.responseBody)

			req := httptest.NewRequest(fiber.MethodDelete, url, nil)
			req.Header.Set(fiber.HeaderAuthorization, fmt.Sprintf("Bearer %s", mockJWTTokenNoExpire))
			resp, _ := mockApp.Test(req)
			defer resp.Body.Close()

			if !tC.isError {
				assert.Equal(t, tC.output.responseStatusCode, resp.StatusCode)
				body, _ := io.ReadAll(resp.Body)
				assert.Equal(t, string(strResponseBodyBytes), strings.TrimSuffix(string(body), "\n"))
			} else {
				assert.Equal(t, tC.output.responseStatusCode, resp.StatusCode)
			}

			tC.after()
		})
	}
}
