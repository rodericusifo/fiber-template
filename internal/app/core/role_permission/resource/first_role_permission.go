package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func (r *RolePermissionResource) FirstRolePermission(query *pkg_types.QuerySQL) (*sql.RolePermission, error) {
	return r.RolePermissionDatabaseSQLRepository.FirstRolePermission(query)
}
