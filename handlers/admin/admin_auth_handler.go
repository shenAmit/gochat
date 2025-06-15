package admin

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/shenAmit/gochat/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginPage(c *fiber.Ctx) error {
	return c.Render("content/authentication/sign-in", fiber.Map{
		"title": "Login - Chat",
	})
}

func LoginHandler(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `form:"email"`
		Password string `form:"password"`
		Remember string `form:"remember"`
	}

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	user := &models.User{}
	err := mgm.Coll(user).First(bson.M{"email": input.Email}, user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// TODO: Replace with secure password check (e.g. bcrypt)
	if user.Password != input.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Incorrect password",
		})
	}

	// Set a cookie as session
	c.Cookie(&fiber.Cookie{
		Name:     "admin_session",
		Value:    user.ID.Hex(),
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Path:     "/",
	})

	// If using HTMX, you may want to swap content or redirect with JS
	return c.SendString(`<script>window.location.href="/";</script>`)
}
