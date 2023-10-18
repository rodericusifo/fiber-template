package constant

import (
	"time"
)

type Default any
type DefaultEnv any

var (
	DEFAULT_TIME_LAYOUT = Default(time.DateTime)
)
var (
	DEFAULT_ENV_SERVER_PORT = DefaultEnv(8080)
)
