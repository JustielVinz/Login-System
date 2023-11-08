package account

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store    *session.Store
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)

func Setup() {

	store = session.New(session.Config{
		CookieHTTPOnly: true,
		// for https
		CookieSecure: true,
		Expiration:   time.Second * 10,
	})

}

func NewMiddleware() fiber.Handler {
	return AuthMiddleware

}

func AuthMiddleware(c *fiber.Ctx) error {

	sess, err := store.Get(c)

	// Simulate obtaining a session token during login
	if strings.Split(c.Path(), "/")[1] == "auth" {
		return c.Next()
	}
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}
	if sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})

	}

	return c.Next()
}
