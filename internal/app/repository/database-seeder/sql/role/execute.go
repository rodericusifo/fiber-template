package role

import (
	"github.com/rodericusifo/fiber-template/internal/pkg/config"

	gorm_seeder "github.com/kachit/gorm-seeder"
	log "github.com/sirupsen/logrus"
)

func ExecuteMysqlRoleDatabaseSeederRepository(isRebuildData config.IsRebuildDataDBSeederMysqlRole, db config.MysqlDatabaseSQLConnection) {
	mysqlRoleDatabaseSeederSQLRepository := InitMysqlRoleDatabaseSeederSQLRepository(gorm_seeder.SeederConfiguration{})
	seedersStack := gorm_seeder.NewSeedersStack(db)
	seedersStack.AddSeeder(mysqlRoleDatabaseSeederSQLRepository)

	if isRebuildData {
		err := seedersStack.Clear()
		if err != nil {
			log.WithFields(log.Fields{
				"message": "clear role fail",
				"detail":  err,
			}).Errorln("[EXECUTE MYSQL ROLE DATABASE SEEDER REPOSITORY]")
			return
		}
		log.WithFields(log.Fields{
			"message": "clear role success",
		}).Infoln("[EXECUTE MYSQL ROLE DATABASE SEEDER REPOSITORY]")
	}

	err := seedersStack.Seed()
	if err != nil {
		log.WithFields(log.Fields{
			"message": "seed role fail",
			"detail":  err,
		}).Errorln("[EXECUTE MYSQL ROLE DATABASE SEEDER REPOSITORY]")
		return
	}
	log.WithFields(log.Fields{
		"message": "seed role success",
	}).Infoln("[EXECUTE MYSQL ROLE DATABASE SEEDER REPOSITORY]")
}
