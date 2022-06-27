package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	AuthModels "github.com/pegdwende/VSM.git/auth/models"
	AuthRoutes "github.com/pegdwende/VSM.git/auth/routes"
	database "github.com/pegdwende/VSM.git/database"
)

func main() {

	db := database.GetConnection()

	db.AutoMigrate(AuthModels.AuthModels...)

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	AuthRoutes.Setup(app)

	fmt.Printf("Made it here")

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	log.Fatal(app.Listen(":4000"))
}
