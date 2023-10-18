package role

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	gorm_seeder "github.com/kachit/gorm-seeder"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
)

type RoleDatabaseSeederSQLRepository struct {
	gorm_seeder.SeederAbstract
	model   sql.Role
	dialect pkg_constant.DialectDatabaseSQL
}

func InitMysqlRoleDatabaseSeederSQLRepository(cfg gorm_seeder.SeederConfiguration) *RoleDatabaseSeederSQLRepository {
	return &RoleDatabaseSeederSQLRepository{gorm_seeder.NewSeederAbstract(cfg), sql.Role{}, pkg_constant.MYSQL}
}
