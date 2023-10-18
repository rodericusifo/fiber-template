package user

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	gorm_seeder "github.com/kachit/gorm-seeder"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
)

type UserDatabaseSeederSQLRepository struct {
	gorm_seeder.SeederAbstract
	models struct {
		sql.User
		sql.Role
	}
	dialect pkg_constant.DialectDatabaseSQL
}

func InitMysqlUserDatabaseSeederSQLRepository(cfg gorm_seeder.SeederConfiguration) *UserDatabaseSeederSQLRepository {
	return &UserDatabaseSeederSQLRepository{gorm_seeder.NewSeederAbstract(cfg), struct{sql.User; sql.Role}{}, pkg_constant.MYSQL}
}
