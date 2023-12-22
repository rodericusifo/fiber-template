package getter

import (
	"github.com/rodericusifo/fiber-template/internal/pkg/config"
)

func GetEnvConfig() config.EnvConfig {
	return config.Env
}
