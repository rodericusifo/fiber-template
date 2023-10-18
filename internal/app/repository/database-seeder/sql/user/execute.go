package user

import (
	"github.com/rodericusifo/fiber-template/internal/pkg/config"

	gorm_seeder "github.com/kachit/gorm-seeder"
	log "github.com/sirupsen/logrus"
)

func ExecuteMysqlUserDatabaseSeederRepository(isRebuildData config.IsRebuildDataDBSeederMysqlUser, db config.MysqlDatabaseSQLConnection) {
	mysqlUserDatabaseSeederSQLRepository := InitMysqlUserDatabaseSeederSQLRepository(gorm_seeder.SeederConfiguration{})
	seedersStack := gorm_seeder.NewSeedersStack(db)
	seedersStack.AddSeeder(mysqlUserDatabaseSeederSQLRepository)

	if isRebuildData {
		err := seedersStack.Clear()
		if err != nil {
			log.WithFields(log.Fields{
				"message": "clear user fail",
				"detail":  err,
			}).Errorln("[EXECUTE MYSQL USER DATABASE SEEDER REPOSITORY]")
			return
		}
		log.WithFields(log.Fields{
			"message": "clear user success",
		}).Infoln("[EXECUTE MYSQL USER DATABASE SEEDER REPOSITORY]")
	}

	err := seedersStack.Seed()
	if err != nil {
		log.WithFields(log.Fields{
			"message": "seed user fail",
			"detail":  err,
		}).Errorln("[EXECUTE MYSQL USER DATABASE SEEDER REPOSITORY]")
		return
	}
	log.WithFields(log.Fields{
		"message": "seed user success",
	}).Infoln("[EXECUTE MYSQL USER DATABASE SEEDER REPOSITORY]")
}
