package account

import (
	"log"
	querywarehouse "sample/middleware/queryWarehouse"
	"sample/middleware/utils/database"
	struct_test "sample/struct"

	"github.com/gofiber/fiber/v2"
)

func FetchData(c *fiber.Ctx) error {
	query := querywarehouse.GetData
	// Fetch data using GORM's Raw method
	var students []struct_test.Student
	result := database.DBConn.Debug().Raw(query).Scan(&students)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return c.JSON(students)
}
