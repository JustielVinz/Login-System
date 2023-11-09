package main

import (
	"fmt"
	"log"
	"os"
	"sample/middleware"
	"sample/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @Title			Log In System
// @Version		1.16.2
// @Description	Log in Code for Go session
// @BasePath		/auth
func main() {

	middleware.CreateConnection()
	// Create a new directory to store the log files, if it does not exist.
	logsDir := "./logs"
	err := os.MkdirAll(logsDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	// Generate a dynamic filename for the log file, based on the current date.
	logFilename := fmt.Sprintf("%s/test %s.log", logsDir, time.Now().Format("2006-01-02"))

	// Open the log file in append mode.
	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logfile.Close()

	// Set the output of the log to the file.
	log.SetOutput(logfile)
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
	}))
	app.Use(recover.New())
	app.Static("/", middleware.GetEnv("STATIC_WEB_LOCATION"))
	app.Use(logger.New())
	routes.SetupRoutes(app)
	// router.SetupPrivateRoutes(app)

	if middleware.GetEnv("SSL") == "enabled" {
		log.Fatal(app.ListenTLS(
			fmt.Sprintf(":%s", middleware.GetEnv("PORT")),
			middleware.GetEnv("SSL_CERTIFICATE"),
			middleware.GetEnv("SSL_KEY"),
		))
	} else {
		err := app.Listen(fmt.Sprintf("%s:%s", middleware.GetEnv("HOST"), middleware.GetEnv("PORT")))
		if err != nil {
			log.Fatal(err.Error())
		}
	}

}
