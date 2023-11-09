package account

import (
	"encoding/json"
	"fmt"
	"sample/middleware/utils/database"
	struct_test "sample/struct"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
// @Router			/signup [post]
// SetupAccount handles the creation of a new user account.
var teacherRegistry []struct_test.AdminAcc

func SetupAccount(c *fiber.Ctx) error {
	newTeacher := struct_test.AdminAcc{}
	successCreating := struct_test.AdminAcc{}
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	// Decode the JSON request body into the AdminAcc struct
	err := json.Unmarshal(c.Body(), &newTeacher)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error decoding request body"})
	}

	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newTeacher.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(struct_test.ErrorResponse{
			Error: "Error hashing password",
		})
	}

	// Check if the username already exists in the database
	if err := database.DBConn.Debug().Raw("INSERT INTO teacher (id, name, department, password, created_at) VALUES (?,?,?,?,?)", newTeacher.ID, newTeacher.Name, newTeacher.Department, hashedPassword, formattedTime).Scan(&successCreating).Error; err != nil {
		// Handle unexpected database error
		fmt.Println("Creating Error:", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(struct_test.ErrorResponse{
			Error: "Database error",
		})
	}

	// Assign a unique ID (you may use a more sophisticated ID generation logic)
	newTeacher.ID = len(teacherRegistry) + 1

	// Add the new teacher to the registry
	teacherRegistry = append(teacherRegistry, newTeacher)

	// Respond with the registered teacher
	return c.Status(fiber.StatusCreated).JSON(successCreating)
}
