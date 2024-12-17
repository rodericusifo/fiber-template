package config

import (
	"context"
	"flag"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/rodericusifo/fiber-template/internal/pkg/migration"
	"github.com/rodericusifo/fiber-template/internal/pkg/types"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/handler"

	jwtware "github.com/gofiber/contrib/jwt"
	log "github.com/sirupsen/logrus"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
)

var (
	Env EnvConfig
)

var (
	MysqlDBSQL    MysqlDatabaseSQLConnection
	PostgresDBSQL PostgresDatabaseSQLConnection
)
var (
	RedisDBCache RedisDatabaseCacheConnection
)

var (
	JWTAuth JWTAuthConfig
)

func ConfigureEnv() {
	var (
		environment = flag.String("env", "", "input the environment type")
	)

	flag.Parse()

	path := fmt.Sprintf("./env/%s.application.env", *environment)
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		log.WithFields(log.Fields{
			"message": "read env fail",
			"detail":  err,
		}).Panic("[CONFIGURE ENV]")
	}

	var env EnvConfig
	if err := viper.Unmarshal(&env); err != nil {
		log.WithFields(log.Fields{
			"message": "load env fail",
			"detail":  err,
		}).Panic("[CONFIGURE ENV]")
	}
	log.WithFields(log.Fields{
		"message": "load env success",
	}).Infoln("[CONFIGURE ENV]")

	Env = env
}

func ConfigureDatabaseSQL(dialect pkg_constant.DialectDatabaseSQL) {
	var (
		dbSQLConfig DBSQLConfig
		db          *gorm.DB
		err         error
	)

	switch dialect {
	case pkg_constant.MYSQL:
		dbSQLConfig = DBSQLConfig{
			Host:              Env.DatabaseMysqlHost,
			Port:              Env.DatabaseMysqlPort,
			Name:              Env.DatabaseMysqlName,
			Username:          Env.DatabaseMysqlUsername,
			Password:          Env.DatabaseMysqlPassword,
			ConnectionTimeout: Env.DatabaseMysqlConnectionTimeout,
			MaxIdleConnection: Env.DatabaseMysqlMaxIdleConnection,
			MaxOpenConnection: Env.DatabaseMysqlMaxOpenConnection,
			DebugMode:         Env.DatabaseMysqlDebugMode,
		}
	case pkg_constant.POSTGRES:
		dbSQLConfig = DBSQLConfig{
			Host:              Env.DatabasePostgresHost,
			Port:              Env.DatabasePostgresPort,
			Name:              Env.DatabasePostgresName,
			Username:          Env.DatabasePostgresUsername,
			Password:          Env.DatabasePostgresPassword,
			TimeZone:          Env.DatabasePostgresTimeZone,
			ConnectionTimeout: Env.DatabasePostgresConnectionTimeout,
			MaxIdleConnection: Env.DatabasePostgresMaxIdleConnection,
			MaxOpenConnection: Env.DatabasePostgresMaxOpenConnection,
			DebugMode:         Env.DatabasePostgresDebugMode,
		}
	}

	cfg := &gorm.Config{
		Logger: logger.Default,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	}

	if dbSQLConfig.DebugMode {
		cfg.Logger = logger.Default.LogMode(logger.Info)
	}

	switch dialect {
	case pkg_constant.MYSQL:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbSQLConfig.Username,
			dbSQLConfig.Password,
			dbSQLConfig.Host,
			dbSQLConfig.Port,
			dbSQLConfig.Name)
		db, err = gorm.Open(mysql.Open(dsn), cfg)
		if err != nil {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("connect to database sql %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE SQL]")
		}
		log.WithFields(log.Fields{
			"message": fmt.Sprintf("connect to database sql %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")

		// Auto Migration Models
		db.AutoMigrate(migration.AutoMigrateModelList...)

		sqlDb, err := db.DB()
		if err != nil {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("set up database sql %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE SQL]")
		}
		log.WithFields(log.Fields{
			"message": fmt.Sprintf("set up database sql %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")
		sqlDb.SetConnMaxIdleTime(dbSQLConfig.ConnectionTimeout)
		sqlDb.SetMaxIdleConns(dbSQLConfig.MaxIdleConnection)
		sqlDb.SetMaxOpenConns(dbSQLConfig.MaxOpenConnection)

		MysqlDBSQL = db
	case pkg_constant.POSTGRES:
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			dbSQLConfig.Host,
			dbSQLConfig.Username,
			dbSQLConfig.Password,
			dbSQLConfig.Name,
			dbSQLConfig.Port,
			dbSQLConfig.TimeZone)
		db, err = gorm.Open(postgres.Open(dsn), cfg)
		if err != nil {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("connect to database sql %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE SQL]")
		}
		log.WithFields(log.Fields{
			"message": fmt.Sprintf("connect to database sql %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")

		// Auto Migration Models
		db.AutoMigrate(migration.AutoMigrateModelList...)

		sqlDb, err := db.DB()
		if err != nil {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("set up database sql %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE SQL]")
		}
		log.WithFields(log.Fields{
			"message": fmt.Sprintf("set up database sql %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")
		sqlDb.SetConnMaxIdleTime(dbSQLConfig.ConnectionTimeout)
		sqlDb.SetMaxIdleConns(dbSQLConfig.MaxIdleConnection)
		sqlDb.SetMaxOpenConns(dbSQLConfig.MaxOpenConnection)

		PostgresDBSQL = db
	}
}

func ConfigureDatabaseCache(dialect pkg_constant.DialectDatabaseCache) {
	switch dialect {
	case pkg_constant.REDIS:
		client := redis.NewClient(&redis.Options{
			Addr:     Env.DatabaseCacheRedisAddress,
			Password: Env.DatabaseCacheRedisPassword,
			Username: Env.DatabaseCacheRedisUsername,
			DB:       Env.DatabaseCacheRedisDatabase,
		})
		ctx := context.Background()
		if err := client.Ping(ctx).Err(); err != nil {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("connect to database cache %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE CACHE]")
		}
		log.WithFields(log.Fields{
			"message": fmt.Sprintf("connect to database cache %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE CACHE]")
		RedisDBCache = client
	}
}

func ConfigureAuth() {
	jwtConfig := &jwtware.Config{
		ErrorHandler: handler.HandleHTTPError,
		Claims:       &types.JwtCustomClaims{},
		SigningKey: jwtware.SigningKey{
			Key: []byte(Env.JWTSecretKey),
		},
	}
	JWTAuth = jwtConfig
}

func ConfigureLog() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
		TimestampFormat:        "2006-01-02 15:04:05 MST",
	})
	log.WithFields(log.Fields{
		"message": "setting log success",
	}).Infoln("[CONFIGURE LOG]")
}
