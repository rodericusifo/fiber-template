package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/app/repository/database/sql/employee"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

type IEmployeeResource interface {
	SaveEmployee(payload *sql.Employee) error
	DeleteEmployee(payload *sql.Employee) error
	FindEmployees(query *pkg_types.QuerySQL) ([]*sql.Employee, error)
	FirstEmployee(query *pkg_types.QuerySQL) (*sql.Employee, error)
	CountEmployees(query *pkg_types.QuerySQL) (int64, error)
}

type EmployeeResource struct {
	EmployeeDatabaseSQLRepository employee.IEmployeeDatabaseSQLRepository
}

func InitEmployeeResource(employeeDatabaseSQLRepository employee.IEmployeeDatabaseSQLRepository) IEmployeeResource {
	return &EmployeeResource{
		EmployeeDatabaseSQLRepository: employeeDatabaseSQLRepository,
	}
}
