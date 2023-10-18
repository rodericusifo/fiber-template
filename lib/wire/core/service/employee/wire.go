//go:build wireinject
// +build wireinject

package employee

import (
	"github.com/google/wire"

	internal_pkg_util_getter "github.com/rodericusifo/fiber-template/internal/pkg/util/getter"

	internal_app_core_employee_resource "github.com/rodericusifo/fiber-template/internal/app/core/employee/resource"
	internal_app_core_employee_service "github.com/rodericusifo/fiber-template/internal/app/core/employee/service"
	internal_app_repository_database_sql_employee "github.com/rodericusifo/fiber-template/internal/app/repository/database/sql/employee"
)

func EmployeeService() internal_app_core_employee_service.IEmployeeService {
	wire.Build(
		internal_pkg_util_getter.GetMysqlDatabaseSQLConnection,
		internal_app_repository_database_sql_employee.InitMysqlEmployeeDatabaseSQLRepository,
		internal_app_core_employee_resource.InitEmployeeResource,
		internal_app_core_employee_service.InitEmployeeService,
	)
	return &internal_app_core_employee_service.EmployeeService{}
}
