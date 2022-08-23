package routes

import (
	"github.com/gofiber/fiber/v2"
	// autorization "github.com/pegdwende/VSM.git/auth/midlleware/authorization"
	InventoryControllers "github.com/pegdwende/VSM.git/inventory/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/create-product", InventoryControllers.CreateProduct)
	app.Post("/api/update-product/:product_id", InventoryControllers.UpdateProduct)
}

func SetUpPublicRoutes(app *fiber.App) {

}
