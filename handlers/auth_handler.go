package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/shenAmit/gochat/middleware/rule"
	"github.com/shenAmit/gochat/models"
)

func ApiTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Register(c *fiber.Ctx) error {
	var input rule.RegisterInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if err := input.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	existing := &models.User{}
	if err := mgm.Coll(existing).First(bson.M{"email": input.Email}, existing); err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Email already in use",
		})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}
	role := &models.Role{}
	err = mgm.Coll(role).First(bson.M{"name": "user"}, role)

	if err != nil {
		role = &models.Role{
			Name:   "user",
			Status: true,
		}
		if err := mgm.Coll(role).Create(role); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create admin role",
			})
		}
	}

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Username: input.Username,
		Password: string(hashedPassword),
		Roles:    []primitive.ObjectID{role.ID},
	}

	if err := mgm.Coll(user).Create(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registration successful",
	})
}
