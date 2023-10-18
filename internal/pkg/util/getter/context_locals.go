package getter

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
	"github.com/rodericusifo/fiber-template/internal/pkg/types"
)

func GetRequestUser(c *fiber.Ctx) *types.RequestUser {
	return c.Locals(constant.CONTEXT_KEY_REQUEST_USER).(*types.RequestUser)
}
