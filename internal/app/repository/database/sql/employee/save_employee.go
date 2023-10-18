package employee

import (
	"gorm.io/gorm/clause"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
)

func (r *EmployeeDatabaseSQLRepository) SaveEmployee(payload *sql.Employee) error {
	employee := new(sql.Employee)

	q := r.db

	if payload != nil {
		employee = payload
	}

	if err := q.Table(r.model.TableName()).Omit(clause.Associations).Save(employee).Error; err != nil {
		return err
	}

	return nil
}
