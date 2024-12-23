package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
)

func (r *EmployeeResource) SaveEmployee(payload *sql.Employee) error {
	return r.EmployeeDatabaseSQLRepository.SaveEmployee(payload)
}
