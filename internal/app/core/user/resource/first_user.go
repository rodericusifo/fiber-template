package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func (r *UserResource) FirstUser(query *pkg_types.QuerySQL) (*sql.User, error) {
	return r.UserDatabaseSQLRepository.FirstUser(query)
}
