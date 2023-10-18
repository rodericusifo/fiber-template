package role

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/validator"

	log "github.com/sirupsen/logrus"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
	pkg_util_builder "github.com/rodericusifo/fiber-template/pkg/util/builder"
)

func (r *RoleDatabaseSeederSQLRepository) Seed(db *gorm.DB) error {
	roles := make([]*sql.Role, 0)
	for _, RoleSeed := range RoleSeedData {
		err := validator.ValidatePayload(RoleSeed)
		if err != nil {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("validation failed: role with xid %s", RoleSeed.XID),
				"detail":  err,
			}).Errorln("[ROLE DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}

		role := new(sql.Role)

		q := db

		query := &pkg_types.QuerySQL{
			Selects: []pkg_types.SelectQuerySQLOperation{
				{Field: "id"},
			},
			Searches: [][]pkg_types.SearchQuerySQLOperation{
				{
					{Field: "xid", Operator: "=", Value: RoleSeed.XID},
					{Field: "slug", Operator: "=", Value: RoleSeed.Slug},
				},
			},
		}

		q = pkg_util_builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)

		err = q.Table(r.model.TableName()).First(role).Error

		if err != nil && err != gorm.ErrRecordNotFound {
			log.WithFields(log.Fields{
				"message": "get role fail",
				"detail":  err,
			}).Errorln("[ROLE DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}
		if role.ID != 0 {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("role with xid %s and slug %s already registered", RoleSeed.XID, RoleSeed.Slug),
			}).Errorln("[ROLE DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}

		roles = append(roles, &sql.Role{
			XID:  RoleSeed.XID,
			Name: RoleSeed.Name,
			Slug: RoleSeed.Slug,
		})
	}
	return db.CreateInBatches(roles, len(roles)).Error
}
