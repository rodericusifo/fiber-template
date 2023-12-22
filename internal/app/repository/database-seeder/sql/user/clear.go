package user

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/pkg/constant"

	log "github.com/sirupsen/logrus"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
	pkg_util_builder "github.com/rodericusifo/fiber-template/pkg/util/builder"
)

func (r *UserDatabaseSeederSQLRepository) Clear(db *gorm.DB) error {
	users := make([]*sql.User, 0)

	q := db

	query := &pkg_types.QuerySQL{
		Selects: []pkg_types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]pkg_types.SearchQuerySQLOperation{
			{
				{Field: "role", Operator: "=", Value: constant.ADMIN},
			},
		},
	}

	q = pkg_util_builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)

	if err := q.Table(r.model.TableName()).Find(&users).Error; err != nil {
		return err
	}

	for _, user := range users {
		if err := db.Table(r.model.TableName()).Delete(user).Error; err != nil {
			log.WithFields(log.Fields{
				"message": "delete user fail",
				"detail":  err,
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [CLEAR]")
			continue
		}
	}

	return nil
}
