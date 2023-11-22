// student_account.go
package account

import (
	"encoding/json"
	"sample/function"
	querywarehouse "sample/middleware/queryWarehouse"
	"sample/middleware/utils/database"
	struct_test "sample/struct"
	errors "sample/struct/error"

	"github.com/gofiber/fiber/v2"
)

// SetupStudentAccount godoc
//
//	@Summary		Create a new student account
//	@Description	Create a new student account with the provided data
//	@Tags			Students
//	@Accept			json
//	@Produce		json
//	@Param			student	body		struct_test.Student	true	"Student data to be created"
//	@Success		201		{object}	struct_test.Student
//	@Failure		400		{object}	errors.ErrorModel
//	@Failure		500		{object}	errors.ErrorModel
//	@Router			/students [post]
func SetupStudentAccount(c *fiber.Ctx) error {
	student := querywarehouse.StudentLogin
	newStudent := struct_test.Student{}
	currentTime := function.GetTime

	// Decode the JSON request body into the Student struct
	err := json.Unmarshal(c.Body(), &newStudent)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.ErrorModel{
			Message:   "Error decoding request body",
			IsSuccess: false,
		})
	}
	// Insert student data
	if err := database.DBConn.Debug().Raw(student, newStudent.ID, newStudent.Name, newStudent.StudentID, newStudent.Department, newStudent.Miscellaneous, newStudent.Payment, newStudent.Amount, currentTime).Scan(&newStudent).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.ErrorModel{
			Message:   "Error in inserting student data",
			IsSuccess: false,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newStudent)
}
