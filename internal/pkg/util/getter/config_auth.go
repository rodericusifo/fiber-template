package getter

import (
	"github.com/rodericusifo/fiber-template/internal/pkg/config"
)

func GetJWTAuthConfig() config.JWTAuthConfig {
	return config.JWTAuth
}
