package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	AuthModels "github.com/pegdwende/VSM.git/auth/models"
	AuthRoutes "github.com/pegdwende/VSM.git/auth/routes"
	database "github.com/pegdwende/VSM.git/database"
	"github.com/pegdwende/VSM.git/env"
	InventoryModels "github.com/pegdwende/VSM.git/inventory/models"
	InventoryRoutes "github.com/pegdwende/VSM.git/inventory/routes"
)

// test loging , protected vs unprotected routes.

func test(c *fiber.Ctx) error {
	fmt.Println("middleware works here man")

	return c.Next()
}

func main() {

	db := database.GetConnection()

	db.AutoMigrate(AuthModels.AuthModels...)
	db.AutoMigrate(InventoryModels.InventoryModels...)

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Get("/healthcheck", test, func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	AuthRoutes.SetupPublicRoutes(app)
	InventoryRoutes.SetUpPublicRoutes(app)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:  []byte(env.GetRequiredEnvVariable("JWT_SECRET")),
		TokenLookup: "cookie:jwt",
	}))

	// app.Group()

	AuthRoutes.Setup(app)
	InventoryRoutes.Setup(app)

	log.Fatal(app.Listen(":4000"))
}
