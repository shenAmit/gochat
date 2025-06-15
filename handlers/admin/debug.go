package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shenAmit/gochat/utils"
)

func DebugSession(c *fiber.Ctx) error {
	sess, _ := utils.Store.Get(c)
	adminID := sess.Get("admin_id")

	if adminID == nil {
		return c.SendString("❌ admin_id not found in session")
	}

	return c.SendString("✅ Logged in as admin: " + adminID.(string))
}
