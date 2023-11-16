package account

import (
	"sample/middleware/utils/database"
	"sample/struct/security"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// SecretKey is the key used for signing the JWT
var SecretKey = []byte("your-secret-key")

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// LoginHandler handles user login and generates a JWT
//
//	@Summary		It will generate the key that will use to log-in in swagger
//	@Description	Auto generate secret key
//	@Produce		json
//	@Param			username	formData	string	true	"Username"
//	@Param			password	formData	string	true	"Password"
//	@Tags			Admin
//
//	@Security		JWT-Token
//
//	@Success		200	{object}	security.SecretKey
//	@Success		200	{string}	string	"Successfully accessed"
//	@Failure		400	{object}	errors.ErrorModel
//	@Failure		401	{object}	errors.ErrorModel
//	@Router			/secure/login [post]
func LoginHandler(c *fiber.Ctx) error {
	requestBody := security.Identification{}
	// Parse JSON request body
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Retrieve user from the database based on the username
	var user security.User
	if err := database.DBConn.Debug().Raw("SELECT username, password  FROM credentials WHERE username = ?", requestBody.Username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Compare the provided password with the hashed password from the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(), // Token expires in 24 hours
		},
	})
	// Sign the token with the secret key
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	// Return the generated token
	return c.JSON(fiber.Map{"token": tokenString})

}
