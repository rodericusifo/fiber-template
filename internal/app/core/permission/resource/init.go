package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/app/repository/database/sql/permission"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

type IPermissionResource interface {
	FirstPermission(query *pkg_types.QuerySQL) (*sql.Permission, error)
}

type PermissionResource struct {
	PermissionDatabaseSQLRepository permission.IPermissionDatabaseSQLRepository
}

func InitPermissionResource(permissionDatabaseSQLRepository permission.IPermissionDatabaseSQLRepository) IPermissionResource {
	return &PermissionResource{
		PermissionDatabaseSQLRepository: permissionDatabaseSQLRepository,
	}
}
