package patcher

import (
	"github.com/rodericusifo/fiber-template/internal/pkg/util/generator"
)

var (
	GenerateHashFromPassword   = generator.GenerateHashFromPassword
	GenerateJWTTokenFromClaims = generator.GenerateJWTTokenFromClaims
)
