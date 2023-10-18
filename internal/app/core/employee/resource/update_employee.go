package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
)

func (r *EmployeeResource) UpdateEmployee(payload *sql.Employee) error {
	return r.EmployeeDatabaseSQLRepository.UpdateEmployee(payload)
}
