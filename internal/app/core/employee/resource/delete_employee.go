package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
)

func (r *EmployeeResource) DeleteEmployee(payload *sql.Employee) error {
	return r.EmployeeDatabaseSQLRepository.DeleteEmployee(payload)
}
