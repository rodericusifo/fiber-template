//go:build wireinject
// +build wireinject

package role

import (
	"github.com/google/wire"

	internal_pkg_util_getter "github.com/rodericusifo/fiber-template/internal/pkg/util/getter"
	internal_app_core_role_resource "github.com/rodericusifo/fiber-template/internal/app/core/role/resource"
	internal_app_repository_database_sql_role "github.com/rodericusifo/fiber-template/internal/app/repository/database/sql/role"
)

func RoleResource() internal_app_core_role_resource.IRoleResource {
	wire.Build(
		internal_pkg_util_getter.GetMysqlDatabaseSQLConnection,
		internal_app_repository_database_sql_role.InitMysqlRoleDatabaseSQLRepository,
		internal_app_core_role_resource.InitRoleResource,
	)
	return &internal_app_core_role_resource.RoleResource{}
}
