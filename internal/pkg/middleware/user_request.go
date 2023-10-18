package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
	"github.com/rodericusifo/fiber-template/internal/pkg/types"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/patcher"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/validator"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func HTTPUserRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals(constant.CONTEXT_KEY_USER).(*jwt.Token).Claims
		user, ok := claims.(*types.JwtCustomClaims)
		if !ok {
			return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Sprintf("invalid claims type. correct type: %T", claims))
		}

		userModelRes, err := patcher.UserResource().FirstUser(&pkg_types.QuerySQL{
			Selects: []pkg_types.SelectQuerySQLOperation{
				{Field: "id"},
				{Field: "name"},
				{Field: "email"},
			},
			Searches: [][]pkg_types.SearchQuerySQLOperation{
				{
					{Field: "xid", Operator: "=", Value: user.XID},
				},
			},
			Joins: []pkg_types.JoinQuerySQLOperation{
				{
					Relation: "Role",
					Selects: []pkg_types.SelectJoinQuerySQLOperation{
						{Field: "id"},
						{Field: "slug"},
					},
				},
			},
		})
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return fiber.NewError(fiber.StatusNotFound, "user not found")
			}
			return err
		}

		reqUser := new(types.RequestUser)
		reqUser = &types.RequestUser{
			ID:    userModelRes.ID,
			XID:   user.XID,
			Name:  userModelRes.Name,
			Email: userModelRes.Email,
			Role: types.RequestRole{
				ID:   userModelRes.Role.ID,
				Slug: userModelRes.Role.Slug,
			},
		}
		if err := validator.ValidateRequestUser(reqUser); err != nil {
			return err
		}

		c.Locals(constant.CONTEXT_KEY_REQUEST_USER, reqUser)
		return c.Next()
	}
}
