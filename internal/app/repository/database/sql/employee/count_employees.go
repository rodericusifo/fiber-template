package employee

import (
	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
	pkg_util_builder "github.com/rodericusifo/fiber-template/pkg/util/builder"
)

func (r *EmployeeDatabaseSQLRepository) CountEmployees(query *pkg_types.QuerySQL) (int64, error) {
	count := int64(0)

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

	q = q.Table(r.model.TableName()).Count(&count)

	if err := q.Error; err != nil {
		return count, err
	}

	return count, nil
}
