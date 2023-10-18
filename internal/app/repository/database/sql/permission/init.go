package permission

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/pkg/config"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

type IPermissionDatabaseSQLRepository interface {
	FirstPermission(query *pkg_types.QuerySQL) (*sql.Permission, error)
}

type PermissionDatabaseSQLRepository struct {
	db      *gorm.DB
	model   sql.Permission
	dialect pkg_constant.DialectDatabaseSQL
}

func InitMysqlPermissionDatabaseSQLRepository(db config.MysqlDatabaseSQLConnection) IPermissionDatabaseSQLRepository {
	return &PermissionDatabaseSQLRepository{
		db:      db,
		model:   sql.Permission{},
		dialect: pkg_constant.MYSQL,
	}
}
