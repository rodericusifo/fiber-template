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
	"github.com/rodericusifo/fiber-template/internal/app/core/employee/controller/api/response"
	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/output"
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	net_url "net/url"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
	pkg_util_response "github.com/rodericusifo/fiber-template/pkg/util/response"
)

func init() {
	SetupTestEmployeeHandler()
}

func TestEmployeeHandler_GetEmployees(t *testing.T) {
	type (
		args struct {
			requestQuery request.GetEmployeesRequestQuery
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
				requestQuery: request.GetEmployeesRequestQuery{
					Page: &mockPageMinus,
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
				requestQuery: request.GetEmployeesRequestQuery{},
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
						arg1 *input.GetEmployeesDTO = &input.GetEmployeesDTO{
							UserID: 1,
						}
					)
					var (
						result output.GetEmployeesDTO = nil
						meta   *pkg_types.Meta        = nil
						err    error                  = errors.New("unexpected errors")
					)
					mockEmployeeService.EXPECT().GetEmployees(arg1).Return(result, meta, err).Once()
				}
			},
			after:   func() {},
			isError: true,
		},
		{
			desc: "[SUCCESS]_success_get_employees",
			input: args{
				requestQuery: request.GetEmployeesRequestQuery{},
			},
			output: result{
				responseStatusCode: fiber.StatusOK,
				responseBody: pkg_util_response.ResponseSuccess("get employees success", []*response.EmployeeResponse{
					{
						XID:       mockUUID,
						Name:      "John",
						Email:     "John@gmail.com",
						Address:   &mockAddress,
						Age:       &mockAge,
						Birthday:  &mockBirthdayString,
						CreatedAt: mockDateString,
						UpdatedAt: mockDateString,
					},
				}, &pkg_types.Meta{
					CurrentPage:      1,
					CountDataPerPage: 1,
					TotalData:        1,
					TotalPage:        1,
				}),
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
						arg1 *input.GetEmployeesDTO = &input.GetEmployeesDTO{
							UserID: 1,
						}
					)
					var (
						result output.GetEmployeesDTO = []*output.EmployeeDTO{
							{
								XID:       mockUUID,
								Name:      "John",
								Email:     "John@gmail.com",
								Address:   &mockAddress,
								Age:       &mockAge,
								Birthday:  &mockBirthdayTime,
								CreatedAt: mockDateTime,
								UpdatedAt: mockDateTime,
							},
						}
						meta *pkg_types.Meta = &pkg_types.Meta{
							CurrentPage:      1,
							CountDataPerPage: 1,
							TotalData:        1,
							TotalPage:        1,
						}
						err error = nil
					)
					mockEmployeeService.EXPECT().GetEmployees(arg1).Return(result, meta, err).Once()
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

			url := "/employees/list"

			q := make(net_url.Values)

			limit := tC.input.requestQuery.Limit
			page := tC.input.requestQuery.Page

			if limit != nil {
				q.Set("limit", fmt.Sprint(*limit))
			}
			if page != nil {
				q.Set("page", fmt.Sprint(*page))
			}

			if len(q) != 0 {
				url = fmt.Sprint(url, "?", q.Encode())
			}

			strResponseBodyBytes, _ := json.Marshal(tC.output.responseBody)

			req := httptest.NewRequest(fiber.MethodGet, url, nil)
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
