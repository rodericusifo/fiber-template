package service

import (
	"time"

	"github.com/rodericusifo/fiber-template/internal/pkg/constant"

	lib_mockery_mocks "github.com/rodericusifo/fiber-template/lib/mockery/mocks"
)

var (
	mockUserResource *lib_mockery_mocks.IUserResource
	mockRoleResource *lib_mockery_mocks.IRoleResource
	authService      IAuthService
)

var (
	mockDateTime                                                           time.Time
	mockUUID, mockPassword, mockHashPassword, mockJWTToken, mockDateString string
)

func SetupTestAuthService() {
	mockUserResource = new(lib_mockery_mocks.IUserResource)
	mockRoleResource = new(lib_mockery_mocks.IRoleResource)

	authService = InitAuthService(mockUserResource, mockRoleResource)

	layoutFormat := constant.DEFAULT_TIME_LAYOUT.(string)

	mockDateString = "2015-09-02 08:04:00"
	mockDateTime, _ = time.Parse(layoutFormat, mockDateString)

	mockUUID = "ac0d6ce3-ff02-4024-896b-ea0ceba32182"

	mockHashPassword = "$2y$14$rnbG3JhbftD.iQV0QRf5GeNI/XlI85KF2kzrf4hnOs48cSoqPvsmG"
	mockPassword = "p4ssw0rd"

	mockJWTToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjhlYTc3OGJjLTM5NTgtNGU5Zi04ZmEyLWE4YTlhZDhmMmFiMSIsIm5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AZ21haWwuY29tIiwicm9sZSI6IkFETUlOIiwiZXhwIjoxNjc3MDc5NzgxfQ.bndXk_BggjadIF2Rwluxc-3tPr-ArfWVYTZ5y03wHU8"
}
