package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func (r *RoleResource) FirstRole(query *pkg_types.QuerySQL) (*sql.Role, error) {
	return r.RoleDatabaseSQLRepository.FirstRole(query)
}
