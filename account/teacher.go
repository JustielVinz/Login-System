// teacher_account.go
package account

import (
	"encoding/json"
	"sample/function"
	"sample/middleware/utils/database"
	struct_test "sample/struct"
	errors "sample/struct/error"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// SetupTeacherAccount godoc
//
//	@Summary		Create a new teacher account
//	@Description	Create a new teacher account with the provided data.
//	@Tags			Teachers
//	@Accept			json
//	@Produce		json
//	@Param			teacher	body	struct_test.AdminAcc	true	"Teacher data to be created"
//	@Security		JWT
//	@Success		201	{object}	struct_test.AdminAcc
//	@Failure		400	{object}	errors.ErrorModel
//	@Failure		401	{object}	errors.ErrorModel
//	@Failure		500	{object}	errors.ErrorModel
//	@Router			/register/teacher [post]
//	@Note			To access this endpoint, you must provide a valid JWT token in the "Authorization" header of your request.
func SetupTeacherAccount(c *fiber.Ctx) error {
	newTeacher := struct_test.AdminAcc{}
	currentTime := function.GetTime

	// Decode the JSON request body into the AdminAcc struct
	err := json.Unmarshal(c.Body(), &newTeacher)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.ErrorModel{
			Message:   "Error decoding request body",
			IsSuccess: false,
		})
	}

	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newTeacher.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.ErrorModel{
			Message:   "Error hashing password",
			IsSuccess: false,
		})
	}

	// Insert teacher data
	if err := database.DBConn.Debug().Raw("INSERT INTO teacher (id, name, department, password, created_at) VALUES (?,?,?,?,?)", newTeacher.ID, newTeacher.Name, newTeacher.Department, hashedPassword, currentTime).Scan(&newTeacher).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.ErrorModel{
			Message:   "Error in inserting teacher data",
			IsSuccess: false,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newTeacher)
}
