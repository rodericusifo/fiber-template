package user

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/patcher"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/validator"

	log "github.com/sirupsen/logrus"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
	pkg_util_builder "github.com/rodericusifo/fiber-template/pkg/util/builder"
)

func (r *UserDatabaseSeederSQLRepository) Seed(db *gorm.DB) error {
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
		}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [SEED]")
		return err
	}

	users := make([]*sql.User, 0)
	for _, UserSeed := range UserSeedData {
		err := validator.ValidatePayload(UserSeed)
		if err != nil {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("validation failed: user with xid %s", UserSeed.XID),
				"detail":  err,
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}

		user := new(sql.User)

		q := db

		query := &pkg_types.QuerySQL{
			Selects: []pkg_types.SelectQuerySQLOperation{
				{Field: "id"},
			},
			Searches: [][]pkg_types.SearchQuerySQLOperation{
				{
					{Field: "xid", Operator: "=", Value: UserSeed.XID},
					{Field: "email", Operator: "=", Value: UserSeed.Email},
				},
			},
		}

		q = pkg_util_builder.BuildQuerySQL(r.models.User.TableName(), q, query, r.dialect)

		err = q.Table(r.models.User.TableName()).First(user).Error

		if err != nil && err != gorm.ErrRecordNotFound {
			log.WithFields(log.Fields{
				"message": "get user fail",
				"detail":  err,
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}
		if user.ID != 0 {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("user with xid %s and email %s already registered", UserSeed.XID, UserSeed.Email),
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}

		hashedPassword, err := patcher.GenerateHashFromPassword(UserSeed.Password)
		if err != nil {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("hash password fail: user with xid %s", UserSeed.XID),
				"detail":  err,
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}

		users = append(users, &sql.User{
			XID:      UserSeed.XID,
			Name:     UserSeed.Name,
			Email:    UserSeed.Email,
			Password: hashedPassword,
			RoleID:   role.ID,
		})
	}
	return db.CreateInBatches(users, len(users)).Error
}
