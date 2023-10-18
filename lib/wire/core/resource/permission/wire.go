//go:build wireinject
// +build wireinject

package permission

import (
	"github.com/google/wire"

	internal_pkg_util_getter "github.com/rodericusifo/fiber-template/internal/pkg/util/getter"
	internal_app_core_permission_resource "github.com/rodericusifo/fiber-template/internal/app/core/permission/resource"
	internal_app_repository_database_sql_permission "github.com/rodericusifo/fiber-template/internal/app/repository/database/sql/permission"
)

func PermissionResource() internal_app_core_permission_resource.IPermissionResource {
	wire.Build(
		internal_pkg_util_getter.GetMysqlDatabaseSQLConnection,
		internal_app_repository_database_sql_permission.InitMysqlPermissionDatabaseSQLRepository,
		internal_app_core_permission_resource.InitPermissionResource,
	)
	return &internal_app_core_permission_resource.PermissionResource{}
}
