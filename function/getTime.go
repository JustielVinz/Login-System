package function

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetTime(c *fiber.Ctx) error {
	// Get the current time
	currentTime := time.Now()
	// Format and print the current time
	Exactime := currentTime.Format("2006-01-02 15:04:05")
	return c.SendString(Exactime)
}
