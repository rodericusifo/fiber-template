package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
)

func (r *EmployeeResource) CreateEmployee(payload *sql.Employee) error {
	return r.EmployeeDatabaseSQLRepository.CreateEmployee(payload)
}
