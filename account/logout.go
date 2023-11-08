package account

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	currentTime := time.Now().Format("15:04:05")
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "logged out (no session)",
		})
	}
	err = sess.Destroy()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}
	log.Printf("Successful Log-out")
	log.Println("Log-out Time: ", currentTime)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logged out",
	})

}
