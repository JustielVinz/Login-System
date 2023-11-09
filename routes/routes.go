package routes

import (
	"sample/account"
	_ "sample/docs"
	"sample/middleware/utils/encryption"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	//swagger
	app.Get("/goroutine_docs/*", swagger.HandlerDefault)

	auth := app.Group("/auth")
	auth.Post("/register",account.SetupAccount)
	auth.Post("/signup", account.SetupAccount)
	auth.Post("/encrypt", encryption.EncryptConn)

}
