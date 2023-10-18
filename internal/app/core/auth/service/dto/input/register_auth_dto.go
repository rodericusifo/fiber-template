package input

import (
	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
)

type RegisterAuthDTO struct {
	Name     string
	Email    string
	Password string
	Role     constant.UserRole
}
