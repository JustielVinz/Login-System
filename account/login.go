package account

import (
	"fmt"
	"log"
	"net/http"
	database "sample/middleware/utils/database"
	"sample/secret"
	struct_test "sample/struct"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Login handles user authentication.
// @Summary 登录
// @Description 登录
// @Produce json
// @Param body body controllers.LoginParams true "body参数"
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
// @Router /user/person/login [post]
func Login(c *fiber.Ctx) error {
	// Struct for log in
	var user struct_test.AdminAcc
	var storedUser struct_test.AdminAcc
	// Parse the request body into the 'user' struct
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	// Retrieve the user's information from your database for the given username
	if err := database.DBConn.Debug().Raw("SELECT username, password, email FROM acc WHERE username = ? ", user.Username).First(&storedUser).Error; err != nil {
		fmt.Println("Query Error:", err.Error())
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
	}
	// Compare the entered password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		fmt.Printf("Password comparison failed: %v\n", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	// Generate a session token (UUID)
	sessionToken := secret.GenerateSessionKey()

	sess, sessErr := store.Get(c)
	if sessErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: in Storing the Session " + sessErr.Error(),
		})
	}

	// Store the session token and additional data in the session
	sess.Set("session_token", sessionToken)
	sess.Set("username", storedUser.Username)

	sessErr = sess.Save()
	if sessErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong in saving the Session: " + sessErr.Error(),
		})
	}
	fmt.Println("my Token: ", sessionToken)
	log.Println("-----------------------Fetch Data----------------------")
	log.Println("Successful in Fetching all the Data in Database")
	log.Println("Host: ", c.Hostname())
	log.Println("-----------------------END-----------------------------")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Login successful",
		"token":    sessionToken,
		"username": storedUser.Username,
	})
}
