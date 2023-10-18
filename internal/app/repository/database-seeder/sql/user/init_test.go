package user

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/mocker"

	gorm_seeder "github.com/kachit/gorm-seeder"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
)

var (
	userDatabaseSeederSQLRepository *UserDatabaseSeederSQLRepository
	mockQuery                       sqlmock.Sqlmock
	mockDB                          *gorm.DB
)

var (
	mockDateTime                     time.Time
	mockHashPassword, mockDateString string
)

func SetupTestUserDatabaseSeederSQLRepository() {
	dialect := pkg_constant.MYSQL
	db, mock := mocker.MockDatabaseSQLConnection(dialect)

	userDatabaseSeederSQLRepository = InitMysqlUserDatabaseSeederSQLRepository(gorm_seeder.SeederConfiguration{})
	mockQuery = mock
	mockDB = db

	layoutFormat := constant.DEFAULT_TIME_LAYOUT.(string)

	mockDateString = "2015-09-02 08:04:00"
	mockDateTime, _ = time.Parse(layoutFormat, mockDateString)

	mockHashPassword = "$2y$14$rnbG3JhbftD.iQV0QRf5GeNI/XlI85KF2kzrf4hnOs48cSoqPvsmG"
}
