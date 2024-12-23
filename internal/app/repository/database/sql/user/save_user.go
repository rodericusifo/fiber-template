package user

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
)

func (r *UserDatabaseSQLRepository) SaveUser(payload *sql.User) error {
	user := new(sql.User)

	q := r.db

	if payload != nil {
		user = payload
	}

	if err := q.Table(r.model.TableName()).Save(user).Error; err != nil {
		return err
	}

	return nil
}
