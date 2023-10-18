package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/patcher"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func (s *AuthService) RegisterAuth(payload *input.RegisterAuthDTO) error {
	payload.RoleSlug = "super_admin"

	userModelRes, err := s.UserResource.FirstUser(&pkg_types.QuerySQL{
		Selects: []pkg_types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]pkg_types.SearchQuerySQLOperation{
			{
				{Field: "email", Operator: "=", Value: payload.Email},
			},
		},
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if userModelRes != nil {
		return fiber.NewError(fiber.StatusConflict, "user already registered")
	}

	roleModelRes, err := s.RoleResource.FirstRole(&pkg_types.QuerySQL{
		Selects: []pkg_types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]pkg_types.SearchQuerySQLOperation{
			{
				{Field: "slug", Operator: "=", Value: payload.RoleSlug},
			},
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, "role not found")
		}
		return err
	}

	hashedPassword, err := patcher.GenerateHashFromPassword(payload.Password)
	if err != nil {
		return err
	}

	userModel := &sql.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
		RoleID:   roleModelRes.ID,
	}
	err = s.UserResource.SaveUser(userModel)
	if err != nil {
		return err
	}

	return nil
}
