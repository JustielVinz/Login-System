package account

import (
	"fmt"
	"log"
	database "sample/middleware/utils/database"
	struct_test "sample/struct"

	"github.com/gofiber/fiber/v2"
)

func GetData(c *fiber.Ctx) error {
	var stored []struct_test.AdminAcc
	if fetcherr := database.DBConn.Debug().Raw("SELECT * FROM acc").Scan(&stored).Error; fetcherr != nil {
		fmt.Println("Fetch Error: ", fetcherr.Error())
	}
	log.Println("-----------------------Fetch Data----------------------")
	log.Println("Successful in Fetching all the Data in Database")
	log.Println("Host: ", c.Hostname())
	log.Println(stored)
	log.Println("-----------------------END-----------------------------")
	return c.JSON(stored)
}
