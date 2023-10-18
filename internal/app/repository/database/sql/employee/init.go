package employee

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/pkg/config"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

type IEmployeeDatabaseSQLRepository interface {
	SaveEmployee(payload *sql.Employee) error
	DeleteEmployee(payload *sql.Employee) error
	FindEmployees(query *pkg_types.QuerySQL) ([]*sql.Employee, error)
	FirstEmployee(query *pkg_types.QuerySQL) (*sql.Employee, error)
	CountEmployees(query *pkg_types.QuerySQL) (int64, error)
}

type EmployeeDatabaseSQLRepository struct {
	db      *gorm.DB
	model   sql.Employee
	dialect pkg_constant.DialectDatabaseSQL
}

func InitMysqlEmployeeDatabaseSQLRepository(db config.MysqlDatabaseSQLConnection) IEmployeeDatabaseSQLRepository {
	return &EmployeeDatabaseSQLRepository{
		db:      db,
		model:   sql.Employee{},
		dialect: pkg_constant.MYSQL,
	}
}
