package employee

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/mocker"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
)

var (
	employeeDatabaseSQLRepository IEmployeeDatabaseSQLRepository
	mockQuery                     sqlmock.Sqlmock
)

var (
	mockDateTime, mockBirthdayTime                            time.Time
	mockUUID, mockAddress, mockDateString, mockBirthdayString string
	mockAge                                                   int
)

func SetupTestMysqlEmployeeDatabaseSQLRepository() {
	dialect := pkg_constant.MYSQL
	db, mock := mocker.MockDatabaseSQLConnection(dialect)

	employeeDatabaseSQLRepository = InitMysqlEmployeeDatabaseSQLRepository(db)
	mockQuery = mock

	layoutFormat := constant.DEFAULT_TIME_LAYOUT.(string)

	mockDateString = "2015-09-02 08:04:00"
	mockDateTime, _ = time.Parse(layoutFormat, mockDateString)

	mockBirthdayString = "1999-03-12 00:00:00"
	mockBirthdayTime, _ = time.Parse(layoutFormat, mockBirthdayString)

	mockUUID = "ac0d6ce3-ff02-4024-896b-ea0ceba32182"
	mockAddress = "Street A, City B"
	mockAge = 25

}
