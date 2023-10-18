//go:build wireinject
// +build wireinject

package user

import (
	"github.com/google/wire"

	internal_pkg_util_getter "github.com/rodericusifo/fiber-template/internal/pkg/util/getter"
	internal_app_core_user_resource "github.com/rodericusifo/fiber-template/internal/app/core/user/resource"
	internal_app_repository_database_sql_user "github.com/rodericusifo/fiber-template/internal/app/repository/database/sql/user"
)

func UserResource() internal_app_core_user_resource.IUserResource {
	wire.Build(
		internal_pkg_util_getter.GetMysqlDatabaseSQLConnection,
		internal_app_repository_database_sql_user.InitMysqlUserDatabaseSQLRepository,
		internal_app_core_user_resource.InitUserResource,
	)
	return &internal_app_core_user_resource.UserResource{}
}
