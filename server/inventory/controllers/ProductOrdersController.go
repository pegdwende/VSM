package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	AuthServices "github.com/pegdwende/VSM.git/auth/services"
	"github.com/pegdwende/VSM.git/database"
	"github.com/pegdwende/VSM.git/inventory/models"
	InventoryModels "github.com/pegdwende/VSM.git/inventory/models"
)

func CreateOders(c *fiber.Ctx) error {

	user := AuthServices.GetUserFromContext(c)

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var existingProduct models.Product

	clientId := int(user["client_id"].(float64))

	productId := data["product_id"]
	database.GetConnection().Joins("JOIN clients on products.client_id = clients.id").
		Where("clients.id=?", clientId).
		Where("products.id=?", productId).First(&existingProduct)

	if existingProduct.ID == 0 {
		return c.JSON(fiber.Map{"message": "Product does not exist "})
	}

	order := InventoryModels.Order{
		ClientID: uint(clientId),
		Status:   models.PENDING_ORDER,
		UserID:   uint(user["user_id"].(float64)),
	}

	order.Create()

	quantity, _ := strconv.Atoi(data["quantity"])
	OrderItem := InventoryModels.OrderItems{
		ProductID: existingProduct.ID,
		OrderID:   order.ID,
		Quantity:  uint(quantity),
		ClientID:  uint(clientId),
	}

	OrderItem.Create()

	return c.JSON(fiber.Map{"product_id": existingProduct.ID, "quantity": quantity, "product_name": existingProduct.ProductName})
}
