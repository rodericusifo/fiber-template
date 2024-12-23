package user

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/pkg/config"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

type IUserDatabaseSQLRepository interface {
	SaveUser(payload *sql.User) error
	FirstUser(query *pkg_types.QuerySQL) (*sql.User, error)
}

type UserDatabaseSQLRepository struct {
	db      *gorm.DB
	model   sql.User
	dialect pkg_constant.DialectDatabaseSQL
}

func InitMysqlUserDatabaseSQLRepository(db config.MysqlDatabaseSQLConnection) IUserDatabaseSQLRepository {
	return &UserDatabaseSQLRepository{
		db:      db,
		model:   sql.User{},
		dialect: pkg_constant.MYSQL,
	}
}
