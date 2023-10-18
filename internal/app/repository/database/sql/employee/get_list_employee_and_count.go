package employee

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
	"github.com/rodericusifo/fiber-template/pkg/util/builder"
)

func (r *EmployeeDatabaseSQLRepository) GetListEmployeeAndCount(query *pkg_types.QuerySQL) ([]*sql.Employee, int, error) {
	count := 0
	employees := make([]*sql.Employee, 0)

	q := r.db

	if query != nil {
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)
	}

	q = q.Table(r.model.TableName()).Find(&employees)

	if err := q.Error; err != nil {
		return nil, count, err
	}

	count = int(q.RowsAffected)

	return employees, count, nil
}
