package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/getter"

	pkg_util_checker "github.com/rodericusifo/fiber-template/pkg/util/checker"
)

func HTTPUserRolesPermission(roles ...constant.UserRole) fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqUser := getter.GetRequestUser(c)

		if reqUser.Role == constant.UserRole("") {
			return fiber.NewError(fiber.StatusUnprocessableEntity, "user role not exist")
		}

		if len(roles) > 0 {
			allowed := pkg_util_checker.CheckSliceContain(roles, reqUser.Role)
			if !allowed {
				return fiber.NewError(fiber.StatusUnauthorized, fmt.Sprintf("user role not allowed. current role: %v. allowed roles: %v", reqUser.Role, roles))
			}
		}

		return c.Next()
	}
}
