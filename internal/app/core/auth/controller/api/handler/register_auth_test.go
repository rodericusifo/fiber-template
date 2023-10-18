package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/fiber-template/internal/app/core/auth/controller/api/request"
	"github.com/rodericusifo/fiber-template/internal/app/core/auth/service/dto/input"

	pkg_util_response "github.com/rodericusifo/fiber-template/pkg/util/response"
)

func init() {
	SetupTestAuthHandler()
}

func TestAuthHandler_RegisterAuth(t *testing.T) {
	type (
		args struct {
			requestBody request.RegisterAuthRequestBody
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
				requestBody: request.RegisterAuthRequestBody{
					Name:     "Ifo",
					Password: "abcd1234",
					// RoleSlug: "super_admin",
				},
			},
			output: result{
				responseStatusCode: fiber.StatusBadRequest,
			},
			before:  func() {},
			after:   func() {},
			isError: true,
		},
		{
			desc: "[ERROR]_because_unexpected_error_from_service",
			input: args{
				requestBody: request.RegisterAuthRequestBody{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: "abcd1234",
					// RoleSlug: "super_admin",
				},
			},
			output: result{
				responseStatusCode: fiber.StatusInternalServerError,
			},
			before: func() {
				{
					var (
						arg1 *input.RegisterAuthDTO = &input.RegisterAuthDTO{
							Name:     "Ifo",
							Email:    "Ifo@gmail.com",
							Password: "abcd1234",
							// RoleSlug: "super_admin",
						}
					)
					var (
						err error = errors.New("unexpected errors")
					)
					mockAuthService.EXPECT().RegisterAuth(arg1).Return(err).Once()
				}
			},
			after:   func() {},
			isError: true,
		},
		{
			desc: "[SUCCESS]_success_register_auth",
			input: args{
				requestBody: request.RegisterAuthRequestBody{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: "abcd1234",
					// RoleSlug: "super_admin",
				},
			},
			output: result{
				responseStatusCode: fiber.StatusOK,
				responseBody:       pkg_util_response.ResponseSuccess[any]("auth register success", nil, nil),
			},
			before: func() {
				{
					var (
						arg1 *input.RegisterAuthDTO = &input.RegisterAuthDTO{
							Name:     "Ifo",
							Email:    "Ifo@gmail.com",
							Password: "abcd1234",
							// RoleSlug: "super_admin",
						}
					)
					var (
						err error = nil
					)
					mockAuthService.EXPECT().RegisterAuth(arg1).Return(err).Once()
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

			url := "/auth/register"

			strRequestBodyBytes, _ := json.Marshal(tC.input.requestBody)
			strResponseBodyBytes, _ := json.Marshal(tC.output.responseBody)

			req := httptest.NewRequest(fiber.MethodPost, url, strings.NewReader(string(strRequestBodyBytes)))
			req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
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
