package employee

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/pkg/config"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

type IEmployeeDatabaseSQLRepository interface {
	CreateEmployee(payload *sql.Employee) error
	UpdateEmployee(payload *sql.Employee) error
	DeleteEmployee(payload *sql.Employee) error
	GetListEmployeeAndCount(query *pkg_types.QuerySQL) ([]*sql.Employee, int, error)
	GetEmployee(query *pkg_types.QuerySQL) (*sql.Employee, error)
	CountAllEmployee(query *pkg_types.QuerySQL) (int, error)
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
