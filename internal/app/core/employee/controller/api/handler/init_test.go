package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
	"github.com/rodericusifo/fiber-template/internal/pkg/types"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/handler"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/patcher"

	jwtware "github.com/gofiber/contrib/jwt"

	internal_app_core_permission_resource "github.com/rodericusifo/fiber-template/internal/app/core/permission/resource"
	internal_app_core_role_permission_resource "github.com/rodericusifo/fiber-template/internal/app/core/role_permission/resource"
	internal_app_core_user_resource "github.com/rodericusifo/fiber-template/internal/app/core/user/resource"
	lib_mockery_mocks "github.com/rodericusifo/fiber-template/lib/mockery/mocks"
)

var (
	mockApp                    *fiber.App
	mockEmployeeService        *lib_mockery_mocks.IEmployeeService
	mockUserResource           *lib_mockery_mocks.IUserResource
	mockPermissionResource     *lib_mockery_mocks.IPermissionResource
	mockRolePermissionResource *lib_mockery_mocks.IRolePermissionResource
	employeeHandler            *EmployeeHandler
)

var (
	mockBirthdayTime, mockDateTime                                                                                                      time.Time
	mockAddress, mockUUID, mockUUIDV1, mockJWTTokenNoExpire, mockUserXID, mockBirthdayString, mockDateString, mockBirthdayInvalidString string
	mockAge, mockAgeMinus, mockPage, mockPageMinus                                                                                      int
)

func SetupTestEmployeeHandler() {
	mockApp = fiber.New(fiber.Config{
		ErrorHandler: handler.HandleHTTPError,
	})

	mockUserResource = new(lib_mockery_mocks.IUserResource)
	patcher.UserResource = func() internal_app_core_user_resource.IUserResource {
		return mockUserResource
	}

	mockPermissionResource = new(lib_mockery_mocks.IPermissionResource)
	patcher.PermissionResource = func() internal_app_core_permission_resource.IPermissionResource {
		return mockPermissionResource
	}

	mockRolePermissionResource = new(lib_mockery_mocks.IRolePermissionResource)
	patcher.RolePermissionResource = func() internal_app_core_role_permission_resource.IRolePermissionResource {
		return mockRolePermissionResource
	}

	mockEmployeeService = new(lib_mockery_mocks.IEmployeeService)

	employee := mockApp.Group("/employees")
	employee.Use(jwtware.New(jwtware.Config{
		ErrorHandler: handler.HandleHTTPError,
		Claims:       &types.JwtCustomClaims{},
		SigningKey: jwtware.SigningKey{
			Key: []byte("zpuCswZDSc"),
		},
	}))
	employeeHandler = InitEmployeeHandler(mockEmployeeService)
	employeeHandler.Mount(employee)

	mockPage = 1
	mockPageMinus = -1

	mockAddress = "20196 Morton Drive"
	mockAge = 24
	mockAgeMinus = -24

	layoutFormat := constant.DEFAULT_TIME_LAYOUT.(string)

	mockBirthdayInvalidString = "2023-08-18T18:51:45+07:00"
	mockBirthdayString = "1999-08-02 08:04:00"
	mockBirthdayTime, _ = time.Parse(layoutFormat, mockBirthdayString)

	mockDateString = "2015-09-02 08:04:00"
	mockDateTime, _ = time.Parse(layoutFormat, mockDateString)

	mockUUID = "802e6e3c-fa65-4148-85d9-d0b7211388b1"
	mockUUIDV1 = "faa3b8a8-ee77-11ed-a05b-0242ac120003"
	mockJWTTokenNoExpire = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ4aWQiOiI4ZWE3NzhiYy0zOTU4LTRlOWYtOGZhMi1hOGE5YWQ4ZjJhYjEifQ.WtQPPQ6BR-TZRyG_NZlOSNE4AngU7C74OkmNOddFdAg"
	mockUserXID = "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1"
}
