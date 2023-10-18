//go:build wireinject
// +build wireinject

package role_permission

import (
	"github.com/google/wire"

	internal_pkg_util_getter "github.com/rodericusifo/fiber-template/internal/pkg/util/getter"
	internal_app_core_role_permission_resource "github.com/rodericusifo/fiber-template/internal/app/core/role_permission/resource"
	internal_app_repository_database_sql_role_permission "github.com/rodericusifo/fiber-template/internal/app/repository/database/sql/role_permission"
)

func RolePermissionResource() internal_app_core_role_permission_resource.IRolePermissionResource {
	wire.Build(
		internal_pkg_util_getter.GetMysqlDatabaseSQLConnection,
		internal_app_repository_database_sql_role_permission.InitMysqlRolePermissionDatabaseSQLRepository,
		internal_app_core_role_permission_resource.InitRolePermissionResource,
	)
	return &internal_app_core_role_permission_resource.RolePermissionResource{}
}
