package middleware

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/rodericusifo/fiber-template/internal/pkg/util/getter"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/patcher"
	"github.com/rodericusifo/fiber-template/pkg/types"
)

func HTTPUserRolePermissions() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqUser := getter.GetRequestUser(c)

		if reqUser.Role.Slug == "super_admin" {
			return c.Next()
		}

		permission, err := patcher.PermissionResource().FirstPermission(&types.QuerySQL{
			Selects: []types.SelectQuerySQLOperation{
				{Field: "id"},
			},
			Searches: [][]types.SearchQuerySQLOperation{
				{
					{Field: "path", Operator: "=", Value: c.OriginalURL()},
				},
			},
		})
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return fiber.NewError(fiber.StatusNotFound, "permission not found")
			}
			return err
		}

		_, err = patcher.RolePermissionResource().FirstRolePermission(&types.QuerySQL{
			Selects: []types.SelectQuerySQLOperation{
				{Field: "id"},
			},
			Searches: [][]types.SearchQuerySQLOperation{
				{
					{Field: "role_id", Operator: "=", Value: reqUser.Role.ID},
					{Field: "permission_id", Operator: "=", Value: permission.ID},
				},
			},
		})
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return fiber.NewError(fiber.StatusNotFound, "permission not allowed")
			}
			return err
		}

		return c.Next()
	}
}
