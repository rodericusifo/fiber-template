package getter

import (
	"strconv"

	"github.com/rodericusifo/fiber-template/internal/pkg/config"
	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
)

func GetEnvServerPort() string {
	if config.Env.ServerPort != 0 {
		return strconv.Itoa(config.Env.ServerPort)
	}
	return strconv.Itoa(constant.DEFAULT_ENV_SERVER_PORT.(int))
}
