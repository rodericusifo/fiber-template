package getter

import (
	"github.com/rodericusifo/fiber-template/internal/pkg/config"
)

func GetMysqlDatabaseSQLConnection() config.MysqlDatabaseSQLConnection {
	return config.MysqlDBSQL
}
func GetPostgresDatabaseSQLConnection() config.PostgresDatabaseSQLConnection {
	return config.PostgresDBSQL
}

func GetRedisDatabaseCacheConnection() config.RedisDatabaseCacheConnection {
	return config.RedisDBCache
}
