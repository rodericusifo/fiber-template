package role

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	log "github.com/sirupsen/logrus"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
	pkg_util_builder "github.com/rodericusifo/fiber-template/pkg/util/builder"
)

func (r *RoleDatabaseSeederSQLRepository) Clear(db *gorm.DB) error {
	roleSeedSlugs := make([]string, 0)
	roles := make([]*sql.Role, 0)

	for _, RoleSeed := range RoleSeedData {
		roleSeedSlugs = append(roleSeedSlugs, RoleSeed.Slug)
	}

	q := db

	query := &pkg_types.QuerySQL{
		Selects: []pkg_types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]pkg_types.SearchQuerySQLOperation{
			{
				{Field: "slug", Operator: "IN", Value: roleSeedSlugs},
			},
		},
	}

	q = pkg_util_builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)

	if err := q.Table(r.model.TableName()).Find(&roles).Error; err != nil {
		return err
	}

	for _, role := range roles {
		if err := db.Table(r.model.TableName()).Delete(role).Error; err != nil {
			log.WithFields(log.Fields{
				"message": "delete role fail",
				"detail":  err,
			}).Errorln("[ROLE DATABASE SEEDER SQL REPOSITORY] [CLEAR]")
			continue
		}
	}

	return nil
}
