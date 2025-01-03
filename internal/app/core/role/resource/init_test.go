package resource

import (
	"time"

	"github.com/rodericusifo/fiber-template/internal/pkg/constant"

	lib_mockery_mocks "github.com/rodericusifo/fiber-template/lib/mockery/mocks"
)

var (
	mockRoleDatabaseSQLRepository *lib_mockery_mocks.IRoleDatabaseSQLRepository
	roleResource                  IRoleResource
)

var (
	mockDateTime                               time.Time
	mockUUID, mockHashPassword, mockDateString string
)

func SetupTestRoleResource() {
	mockRoleDatabaseSQLRepository = new(lib_mockery_mocks.IRoleDatabaseSQLRepository)

	roleResource = InitRoleResource(mockRoleDatabaseSQLRepository)

	layoutFormat := constant.DEFAULT_TIME_LAYOUT.(string)

	mockDateString = "2015-09-02 08:04:00"
	mockDateTime, _ = time.Parse(layoutFormat, mockDateString)

	mockUUID = "ac0d6ce3-ff02-4024-896b-ea0ceba32182"
	mockHashPassword = "$2y$14$rnbG3JhbftD.iQV0QRf5GeNI/XlI85KF2kzrf4hnOs48cSoqPvsmG"
}
