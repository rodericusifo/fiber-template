package user

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/mocker"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
)

var (
	userDatabaseSQLRepository IUserDatabaseSQLRepository
	mockQuery                 sqlmock.Sqlmock
)

var (
	mockDateTime                               time.Time
	mockUUID, mockHashPassword, mockDateString string
)

func SetupTestMysqlUserDatabaseSQLRepository() {
	dialect := pkg_constant.MYSQL
	db, mock := mocker.MockDatabaseSQLConnection(dialect)

	userDatabaseSQLRepository = InitMysqlUserDatabaseSQLRepository(db)
	mockQuery = mock

	layoutFormat := constant.DEFAULT_TIME_LAYOUT.(string)

	mockDateString = "2015-09-02 08:04:00"
	mockDateTime, _ = time.Parse(layoutFormat, mockDateString)

	mockUUID = "ac0d6ce3-ff02-4024-896b-ea0ceba32182"
	mockHashPassword = "$2y$14$rnbG3JhbftD.iQV0QRf5GeNI/XlI85KF2kzrf4hnOs48cSoqPvsmG"
}
