package employee

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
	pkg_util_builder "github.com/rodericusifo/fiber-template/pkg/util/builder"
)

func (r *EmployeeDatabaseSQLRepository) CountAllEmployee(query *pkg_types.QuerySQL) (int, error) {
	count := 0
	employees := make([]*sql.Employee, 0)

	q := r.db

	defaultQuery := &pkg_types.QuerySQL{
		Selects: []pkg_types.SelectQuerySQLOperation{
			{Field: "id"},
		},
	}

	if query != nil {
		query.Selects = append(query.Selects, defaultQuery.Selects...)
		q = pkg_util_builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)
	} else {
		q = pkg_util_builder.BuildQuerySQL(r.model.TableName(), q, defaultQuery, r.dialect)
	}

	q = q.Table(r.model.TableName()).Find(&employees)

	if err := q.Error; err != nil {
		return count, err
	}

	count = int(q.RowsAffected)

	return count, nil
}
