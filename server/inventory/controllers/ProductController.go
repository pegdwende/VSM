package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pegdwende/VSM.git/database"
	"github.com/pegdwende/VSM.git/inventory/models"
	InventoryModels "github.com/pegdwende/VSM.git/inventory/models"
)

func CreateProduct(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var existingProduct models.Product

	database.GetConnection().Joins("JOIN client on products.client_id = client.id").
		Where("client.client_code=?", data["client_code"]).
		Where("products.product_name=?", data["product_name"]).First(&existingProduct)

	if existingProduct.ID != 0 {
		return c.JSON(fiber.Map{"message": "Product Name exist "})
	}

	clientId, _ := strconv.Atoi(data["role_id"])
	quantity, _ := strconv.Atoi(data["quantity"])
	price, _ := strconv.ParseFloat(data["price"], 64)
	product := InventoryModels.Product{
		ClientID:           uint(clientId),
		ProductDescription: data["description"],
		ProductImage:       data["image"],
		ProductQuantity:    uint(quantity),
		ProductPrice:       price,
		ProductName:        data["name"],
	}

	product.Create()

	return c.JSON(product)

}
