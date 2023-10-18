package service

import (
	"github.com/rodericusifo/fiber-template/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/app/core/auth/service/dto/output"
	
	internal_app_core_user_resource "github.com/rodericusifo/fiber-template/internal/app/core/user/resource"
	internal_app_core_role_resource "github.com/rodericusifo/fiber-template/internal/app/core/role/resource"
)

type IAuthService interface {
	RegisterAuth(payload *input.RegisterAuthDTO) error
	LoginAuth(payload *input.LoginAuthDTO) (*output.LoginAuthDTO, error)
}

type AuthService struct {
	UserResource internal_app_core_user_resource.IUserResource
	RoleResource internal_app_core_role_resource.IRoleResource
}

func InitAuthService(userResource internal_app_core_user_resource.IUserResource, roleResource internal_app_core_role_resource.IRoleResource) IAuthService {
	return &AuthService{
		UserResource: userResource,
		RoleResource: roleResource,
	}
}
