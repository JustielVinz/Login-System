package encryption

import (
	"sample/algorithm"
	errors "sample/struct/error"
	requestenv "sample/struct/requestEnv"

	"github.com/gofiber/fiber/v2"
)

func EncryptConn(c *fiber.Ctx) error {
	userCrendentials := requestenv.Env{}

	if parsErr := c.BodyParser(&userCrendentials); parsErr != nil {
		return c.JSON(errors.ResponseModel{
			RetCode: "400",
			Message: "Bad Request",
			Data: errors.ErrorModel{
				Message:   "Error: Body Parsing",
				IsSuccess: false,
				Error:     parsErr,
			},
		})
	}
	encryptedCredentials := requestenv.Env{
		DatabaseName:     algorithm.EncodeBase64(userCrendentials.DatabaseName),
		PostgresUsename:  algorithm.EncodeBase64(userCrendentials.PostgresUsename),
		PostgresPassword: algorithm.EncodeBase64(userCrendentials.PostgresPassword),
		PostgresHost:     algorithm.EncodeBase64(userCrendentials.PostgresHost),
		PostgresPort:     algorithm.EncodeBase64(userCrendentials.PostgresPort),
	}

	return c.JSON(errors.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    encryptedCredentials,
	})

}
