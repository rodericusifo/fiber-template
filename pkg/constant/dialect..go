package constant

type DialectDatabaseSQL string
type DialectDatabaseCache string

var (
	POSTGRES = DialectDatabaseSQL("postgres")
	MYSQL    = DialectDatabaseSQL("mysql")
)
var (
	REDIS = DialectDatabaseCache("redis")
)
