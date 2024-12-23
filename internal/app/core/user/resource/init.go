package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/app/repository/database/sql/user"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

type IUserResource interface {
	SaveUser(payload *sql.User) error
	FirstUser(query *pkg_types.QuerySQL) (*sql.User, error)
}

type UserResource struct {
	UserDatabaseSQLRepository user.IUserDatabaseSQLRepository
}

func InitUserResource(userDatabaseSQLRepository user.IUserDatabaseSQLRepository) IUserResource {
	return &UserResource{
		UserDatabaseSQLRepository: userDatabaseSQLRepository,
	}
}
