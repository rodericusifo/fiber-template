package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/app/repository/database/sql/role"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

type IRoleResource interface {
	FirstRole(query *pkg_types.QuerySQL) (*sql.Role, error)
}

type RoleResource struct {
	RoleDatabaseSQLRepository role.IRoleDatabaseSQLRepository
}

func InitRoleResource(roleDatabaseSQLRepository role.IRoleDatabaseSQLRepository) IRoleResource {
	return &RoleResource{
		RoleDatabaseSQLRepository: roleDatabaseSQLRepository,
	}
}
