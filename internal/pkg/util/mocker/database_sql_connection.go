package mocker

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
)

func MockDatabaseSQLConnection(dialect pkg_constant.DialectDatabaseSQL) (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp),
	)
	if err != nil {
		log.WithFields(log.Fields{
			"message": fmt.Sprintf("connect to mock database sql %s failed", dialect),
			"detail":  err,
		}).Panic("[MOCK CONNECTION DATABASE SQL]")
	}

	var dialector gorm.Dialector
	switch dialect {
	case pkg_constant.POSTGRES:
		dialector = postgres.New(postgres.Config{
			Conn:       sqlDB,
			DriverName: string(pkg_constant.POSTGRES),
		})
	case pkg_constant.MYSQL:
		dialector = mysql.New(mysql.Config{
			Conn:                      sqlDB,
			DriverName:                string(pkg_constant.MYSQL),
			SkipInitializeWithVersion: true,
		})
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.WithFields(log.Fields{
			"message": fmt.Sprintf("connect to mock database sql %s failed", dialect),
			"detail":  err,
		}).Panic("[MOCK CONNECTION DATABASE SQL]")
	}

	return db, mock
}
