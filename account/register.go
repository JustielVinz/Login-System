package account

import (
	"fmt"
	"log"
	"sample/hashing"
	database "sample/middleware/utils/database"
	struct_test "sample/struct"
	"time"

	"github.com/gofiber/fiber/v2"
)

// @Summary		Set up the Account
// @Description	The user will create an Account
// @Tags			Creating New Account
// @ID				createAccount
// @Param			request	body	struct_test.AdminAcc	true	"Create  Account"
// @Accept			json
// @Produce		json
// @Success		200	{object}	struct_test.AdminAcc
// @Failure		404	{object}	struct_test.ErrorResponse
// @Router			/auth/signup [post]
// SetupAccount handles the creation of a new user account.
func SetupAccount(c *fiber.Ctx) error {
	currentTime := time.Now().Format("15:04:05")
	c.Accepts("application/json")
	// Parse the request body into the 'acc' struct
	var acc struct_test.AdminAcc
	var dataUser struct_test.AdminAcc
	var existingUser struct_test.AdminAcc
	if err := c.BodyParser(&acc); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(struct_test.ErrorResponse{
			Error: "Invalid request body",
		})
	}
	// Check if the username already exists in the database
	if err := database.DBConn.Debug().Raw("SELECT username FROM acc WHERE username = ?", acc.Username).First(&existingUser).Error; err != nil {
		if err.Error() != "record not found" {
			// Handle unexpected database error
			fmt.Println("Database Error:", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(struct_test.ErrorResponse{
				Error: "Database error",
			})
		}
	}
	if existingUser.Username != "" {
		// Username already exists
		return c.Status(fiber.StatusConflict).JSON(struct_test.ErrorResponse{
			Error: "Username already in use",
		})
	}
	// Hash the password
	hashedPassword, err := hashing.HashPassword(acc.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(struct_test.ErrorResponse{
			Error: "Failed to hash the password",
		})
	}

	acc.Password = string(hashedPassword)
	// Insert account information into the database
	if err := database.DBConn.Debug().Raw("INSERT INTO acc (id, name , username , password, role) VALUES (?,?,?,?,?)", acc.ID, acc.Name, acc.Username, acc.Password, acc.Role).Scan(&dataUser).Error; err != nil {
		fmt.Println("Insert Error: ", err.Error())
		return c.JSON(struct_test.ErrorResponse{
			Error: err.Error(),
		})
	}
	log.Println("Success Creaing Account")
	log.Println("Time: ", currentTime)
	return c.SendString("Register")
}
