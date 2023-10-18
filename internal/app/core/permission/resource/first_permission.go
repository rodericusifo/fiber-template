package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func (r *PermissionResource) FirstPermission(query *pkg_types.QuerySQL) (*sql.Permission, error) {
	return r.PermissionDatabaseSQLRepository.FirstPermission(query)
}
