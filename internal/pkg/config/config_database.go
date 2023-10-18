package config

import (
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type (
	MysqlDatabaseSQLConnection    *gorm.DB
	PostgresDatabaseSQLConnection *gorm.DB
)

type (
	RedisDatabaseCacheConnection *redis.Client
)

type DBSQLConfig struct {
	Host              string
	Port              string
	Name              string
	Username          string
	Password          string
	TimeZone          string
	ConnectionTimeout time.Duration
	MaxIdleConnection int
	MaxOpenConnection int
	DebugMode         bool
}
