// teacher_account.go
package account

import (
	"sample/middleware/utils/database"
	struct_test "sample/struct"
	errors "sample/struct/error"
	"time"

	"github.com/gofiber/fiber/v2"
)

// SetupTeacherAccount

// @Summary		Create a new teacher
// @Description	Add a new teacher to the database
// @Tags			teachers
// @Accept			json
// @Produce		json
// @Param			department	body		struct_test.AdminAcc	true	"Enter the Description"
// @Success		200			{object}	string					"Teacher added successfully"
// @Failure		400			{object}	errors.ErrorModel
//
// @Router			/register/teacher [post]
// @Security		JWT-Token
// @Note			To access this endpoint, you must provide a valid JWT token in the "Authorization" header of your request.
func SetupTeacherAccount(c *fiber.Ctx) error {
	newTeacher := struct_test.AdminAcc{}
	// Get the current time
	currentTime := time.Now()
	// Format and print the current time
	Exactime := currentTime.Format("2006-01-02 15:04:05")
	if err := c.BodyParser(&newTeacher); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	// Insert teacher data
	if err := database.DBConn.Debug().Raw("INSERT INTO teacher (name, department, created_at) VALUES (?,?,?)", newTeacher.Name, newTeacher.Department, Exactime).Scan(&newTeacher).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.ErrorModel{
			Message:   "Error in inserting teacher data",
			IsSuccess: false,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newTeacher)
}
