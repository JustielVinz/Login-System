package routes

import (
	"sample/account"
	_ "sample/docs"
	"sample/function"
	"sample/middleware/utils/encryption"
	"sample/secret"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	//swagger
	app.Get("/payment_docs/*", swagger.HandlerDefault)

	auth := app.Group("/auth")
	auth.Post("/encrypt", encryption.EncryptConn)
	auth.Post("/register/teacher", account.SetupTeacherAccount)
	auth.Post("/register/student", account.SetupStudentAccount)
	auth.Get("/protected", secret.Authenticate, function.ProtectedEndpoint)

}
