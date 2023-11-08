package account

import (
	"fmt"
	"log"
	database "sample/middleware/utils/database"
	struct_test "sample/struct"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ParamData(c *fiber.Ctx) error {
	currentTime := time.Now().Format("15:04:05")
	model := struct_test.AdminAcc{}
	id := c.Params("id")
	if getErr := database.DBConn.Debug().Raw("SELECT * FROM acc WHERE id = ?", id).Scan(&model).Error; getErr != nil {
		fmt.Println("Fetch Data Error: ", getErr.Error())
	}
	model.Password = " "
	log.Println("Successful in Fetching all the Data in Database")
	log.Println("Time: ", currentTime)
	return c.JSON(model)
}
