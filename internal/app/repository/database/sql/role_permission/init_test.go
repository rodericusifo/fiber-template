package role_permission

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/mocker"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
)

var (
	rolePermissionDatabaseSQLRepository IRolePermissionDatabaseSQLRepository
	mockQuery                 sqlmock.Sqlmock
)

var (
	mockDateTime             time.Time
	mockUUID, mockDateString string
)

func SetupTestMysqlRolePermissionDatabaseSQLRepository() {
	dialect := pkg_constant.MYSQL
	db, mock := mocker.MockDatabaseSQLConnection(dialect)

	rolePermissionDatabaseSQLRepository = InitMysqlRolePermissionDatabaseSQLRepository(db)
	mockQuery = mock

	layoutFormat := constant.DEFAULT_TIME_LAYOUT.(string)

	mockDateString = "2015-09-02 08:04:00"
	mockDateTime, _ = time.Parse(layoutFormat, mockDateString)

	mockUUID = "ac0d6ce3-ff02-4024-896b-ea0ceba32182"
}
