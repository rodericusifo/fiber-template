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
	userModelRes, err := s.UserResource.GetUser(&pkg_types.QuerySQL{
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

	hashedPassword, err := patcher.GenerateHashFromPassword(payload.Password)
	if err != nil {
		return err
	}

	userModel := &sql.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
		Role:     payload.Role,
	}
	err = s.UserResource.CreateUser(userModel)
	if err != nil {
		return err
	}

	return nil
}
