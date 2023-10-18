package permission

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
	pkg_util_builder "github.com/rodericusifo/fiber-template/pkg/util/builder"
)

func (r *PermissionDatabaseSQLRepository) FirstPermission(query *pkg_types.QuerySQL) (*sql.Permission, error) {
	permission := new(sql.Permission)

	q := r.db

	if query != nil {
		q = pkg_util_builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)
	}

	if err := q.Table(r.model.TableName()).First(permission).Error; err != nil {
		return nil, err
	}

	return permission, nil
}
