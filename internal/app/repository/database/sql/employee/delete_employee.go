package employee

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
)

func (r *EmployeeDatabaseSQLRepository) DeleteEmployee(payload *sql.Employee) error {
	employee := new(sql.Employee)

	q := r.db

	if payload != nil {
		employee = payload
	}

	if err := q.Table(r.model.TableName()).Unscoped().Delete(employee).Error; err != nil {
		return err
	}

	return nil
}
