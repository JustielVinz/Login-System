package function

import "github.com/gofiber/fiber/v2"

func ProtectedEndpoint(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "This is a protected endpoint",
	})
}
