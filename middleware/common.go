package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shenAmit/gochat/utils"
)

func IsAdmin(c *fiber.Ctx) error {
	sess, _ := utils.Store.Get(c)
	adminID := sess.Get("admin_id")
	if adminID == nil {
		return c.Redirect("/v1/login")
	}

	c.Locals("admin_id", adminID)
	return c.Next()
}
