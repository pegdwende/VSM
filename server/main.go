package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	AuthModels "github.com/pegdwende/VSM.git/auth/models"
	AuthRoutes "github.com/pegdwende/VSM.git/auth/routes"
	database "github.com/pegdwende/VSM.git/database"
	"github.com/pegdwende/VSM.git/env"
)

// test loging , protected vs unprotected routes.

func main() {

	db := database.GetConnection()

	db.AutoMigrate(AuthModels.AuthModels...)

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:  []byte(env.GetRequiredEnvVariable("JWT_SECRET")),
		TokenLookup: "cookie:jwt",
	}))

	AuthRoutes.Setup(app)

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	log.Fatal(app.Listen(":4000"))
}
