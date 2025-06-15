package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"github.com/shenAmit/gochat/models"
	"github.com/shenAmit/gochat/utils"
)

func LoginPage(c *fiber.Ctx) error {
	sess, _ := utils.Store.Get(c)
	adminID := sess.Get("admin_id")

	if adminID != nil {
		return c.Redirect("/v1")
	}
	return c.Render("content/authentication/sign-in", fiber.Map{
		"title": "Login - Chat",
	})
}

func LoginHandler(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	user := &models.User{}
	if err := mgm.Coll(user).First(bson.M{"email": input.Email}, user); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Incorrect password",
		})
	}

	sess, _ := utils.Store.Get(c)
	sess.Set("admin_id", user.ID.Hex())
	sess.Save()

	return c.SendString(`<script>window.location.href="/v1";</script>`)
}

func LogoutHandler(c *fiber.Ctx) error {
	sess, _ := utils.Store.Get(c)
	sess.Destroy()
	sess.Save()

	return c.Redirect("/v1/login")
}
