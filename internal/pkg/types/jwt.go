package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	XID   string `json:"xid"`
	jwt.RegisteredClaims
}
