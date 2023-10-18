package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/app/repository/database/sql/employee"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

type IEmployeeResource interface {
	CreateEmployee(payload *sql.Employee) error
	UpdateEmployee(payload *sql.Employee) error
	DeleteEmployee(payload *sql.Employee) error
	GetListEmployeeAndCount(query *pkg_types.QuerySQL) ([]*sql.Employee, int, error)
	GetEmployee(query *pkg_types.QuerySQL) (*sql.Employee, error)
	CountAllEmployee(query *pkg_types.QuerySQL) (int, error)
}

type EmployeeResource struct {
	EmployeeDatabaseSQLRepository employee.IEmployeeDatabaseSQLRepository
}

func InitEmployeeResource(employeeDatabaseSQLRepository employee.IEmployeeDatabaseSQLRepository) IEmployeeResource {
	return &EmployeeResource{
		EmployeeDatabaseSQLRepository: employeeDatabaseSQLRepository,
	}
}
