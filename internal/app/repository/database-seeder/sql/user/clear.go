package user

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	log "github.com/sirupsen/logrus"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
	pkg_util_builder "github.com/rodericusifo/fiber-template/pkg/util/builder"
)

func (r *UserDatabaseSeederSQLRepository) Clear(db *gorm.DB) error {
	role := new(sql.Role)

	q := db

	query := &pkg_types.QuerySQL{
		Selects: []pkg_types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]pkg_types.SearchQuerySQLOperation{
			{
				{Field: "slug", Operator: "=", Value: "super_admin"},
			},
		},
	}

	q = pkg_util_builder.BuildQuerySQL(r.models.Role.TableName(), q, query, r.dialect)

	err := q.Table(r.models.Role.TableName()).First(role).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"message": "get role fail",
			"detail":  err,
		}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [CLEAR]")
		return err
	}

	users := make([]*sql.User, 0)

	q = db

	query = &pkg_types.QuerySQL{
		Selects: []pkg_types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]pkg_types.SearchQuerySQLOperation{
			{
				{Field: "role_id", Operator: "=", Value: role.ID},
			},
		},
	}

	q = pkg_util_builder.BuildQuerySQL(r.models.User.TableName(), q, query, r.dialect)

	if err := q.Table(r.models.User.TableName()).Find(&users).Error; err != nil {
		return err
	}

	for _, user := range users {
		if err := db.Table(r.models.User.TableName()).Delete(user).Error; err != nil {
			log.WithFields(log.Fields{
				"message": "delete user fail",
				"detail":  err,
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [CLEAR]")
			continue
		}
	}

	return nil
}
