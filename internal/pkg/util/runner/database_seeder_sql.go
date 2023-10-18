package runner

import (
	"github.com/rodericusifo/fiber-template/internal/app/repository/database-seeder/sql/user"
	"github.com/rodericusifo/fiber-template/internal/app/repository/database-seeder/sql/role"
	"github.com/rodericusifo/fiber-template/internal/pkg/config"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/getter"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
)

func RunDatabaseSeederSQL(dialect pkg_constant.DialectDatabaseSQL) {
	switch dialect {
	case pkg_constant.POSTGRES:
	case pkg_constant.MYSQL:
		role.ExecuteMysqlRoleDatabaseSeederRepository(config.Env.DatabaseSeederMysqlRoleIsRebuildData, getter.GetMysqlDatabaseSQLConnection())
		user.ExecuteMysqlUserDatabaseSeederRepository(config.Env.DatabaseSeederMysqlUserIsRebuildData, getter.GetMysqlDatabaseSQLConnection())
	}
}
