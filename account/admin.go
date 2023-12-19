package account

import (
	querywarehouse "sample/middleware/queryWarehouse"
	"sample/middleware/utils/database"
	errors "sample/struct/error"
	"sample/struct/security"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *fiber.Ctx) error {
	keyAdmin := querywarehouse.AdminKey
	admin := security.User{}
	// Get the current time
	currentTime := time.Now()
	// Format and print the current time
	Exactime := currentTime.Format("2006-01-02 15:04:05")

	// Parse JSON request body
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	// Insert the user into the database
	if err := database.DBConn.Debug().Raw(keyAdmin, admin.ID, admin.Username, hashedPassword, admin.Status, Exactime).Scan(&admin).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.ErrorModel{
			Message:   "Error in inserting user data",
			IsSuccess: false,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}
