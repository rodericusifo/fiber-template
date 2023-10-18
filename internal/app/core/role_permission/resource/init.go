package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/app/repository/database/sql/role_permission"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

type IRolePermissionResource interface {
	FirstRolePermission(query *pkg_types.QuerySQL) (*sql.RolePermission, error)
}

type RolePermissionResource struct {
	RolePermissionDatabaseSQLRepository role_permission.IRolePermissionDatabaseSQLRepository
}

func InitRolePermissionResource(rolePermissionDatabaseSQLRepository role_permission.IRolePermissionDatabaseSQLRepository) IRolePermissionResource {
	return &RolePermissionResource{
		RolePermissionDatabaseSQLRepository: rolePermissionDatabaseSQLRepository,
	}
}
