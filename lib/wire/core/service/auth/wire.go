//go:build wireinject
// +build wireinject

package auth

import (
	"github.com/google/wire"

	internal_pkg_util_getter "github.com/rodericusifo/fiber-template/internal/pkg/util/getter"

	internal_app_core_auth_service "github.com/rodericusifo/fiber-template/internal/app/core/auth/service"
	internal_app_core_user_resource "github.com/rodericusifo/fiber-template/internal/app/core/user/resource"
	internal_app_core_role_resource "github.com/rodericusifo/fiber-template/internal/app/core/role/resource"
	internal_app_repository_database_sql_user "github.com/rodericusifo/fiber-template/internal/app/repository/database/sql/user"
	internal_app_repository_database_sql_role "github.com/rodericusifo/fiber-template/internal/app/repository/database/sql/role"
)

func AuthService() internal_app_core_auth_service.IAuthService {
	wire.Build(
		internal_pkg_util_getter.GetMysqlDatabaseSQLConnection,
		internal_app_repository_database_sql_user.InitMysqlUserDatabaseSQLRepository,
		internal_app_repository_database_sql_role.InitMysqlRoleDatabaseSQLRepository,
		internal_app_core_user_resource.InitUserResource,
		internal_app_core_role_resource.InitRoleResource,
		internal_app_core_auth_service.InitAuthService,
	)
	return &internal_app_core_auth_service.AuthService{}
}
