package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
)

func (r *UserResource) CreateUser(payload *sql.User) error {
	return r.UserDatabaseSQLRepository.CreateUser(payload)
}
