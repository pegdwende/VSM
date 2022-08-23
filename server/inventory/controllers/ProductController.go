package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	AuthServices "github.com/pegdwende/VSM.git/auth/services"
	"github.com/pegdwende/VSM.git/database"
	"github.com/pegdwende/VSM.git/inventory/models"
	InventoryModels "github.com/pegdwende/VSM.git/inventory/models"
)

func CreateProduct(c *fiber.Ctx) error {

	user := AuthServices.GetUserFromContext(c)

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var existingProduct models.Product

	fmt.Println(user)
	fmt.Println(user["client_id"])

	clientId := int(user["client_id"].(float64))

	database.GetConnection().Joins("JOIN clients on products.client_id = clients.id").
		Where("clients.id=?", clientId).
		Where("products.product_name=?", data["product_name"]).First(&existingProduct)

	if existingProduct.ID != 0 {
		return c.JSON(fiber.Map{"message": "Product Name exist "})
	}

	quantity, _ := strconv.Atoi(data["quantity"])
	price, _ := strconv.ParseFloat(data["price"], 64)
	product := InventoryModels.Product{
		ClientID:           uint(clientId),
		ProductDescription: data["description"],
		ProductImage:       data["image"],
		ProductQuantity:    uint(quantity),
		ProductPrice:       price,
		ProductName:        data["product_name"],
	}

	product.Create()

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {

	user := AuthServices.GetUserFromContext(c)

	var existingProduct models.Product

	clientId := int(user["client_id"].(float64))

	product_id := c.Params("product_id", "0")

	database.GetConnection().Joins("JOIN clients on products.client_id = clients.id").
		Where("clients.id=?", clientId).
		Where("products.id=?", product_id).First(&existingProduct)

	if existingProduct.ID == 0 {
		return c.JSON(fiber.Map{"message": "Product does does not exist for this client"})
	}

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	quantity := existingProduct.ProductQuantity

	if val, ok := data["quantity"]; ok {
		quantityInt, _ := strconv.Atoi(val)
		quantity = uint(quantityInt)
	}

	description := existingProduct.ProductDescription

	if val, ok := data["description"]; ok {
		description = val
	}

	productName := existingProduct.ProductName

	if val, ok := data["product_name"]; ok {
		productName = val
	}

	imageUrl := existingProduct.ProductImage

	if val, ok := data["image"]; ok {
		imageUrl = val
	}

	price := existingProduct.ProductPrice

	if val, ok := data["price"]; ok {
		priceFloat, _ := strconv.ParseFloat(val, 64)
		price = priceFloat
	}

	existingProduct.ClientID = uint(clientId)
	existingProduct.ProductDescription = description
	existingProduct.ProductImage = imageUrl
	existingProduct.ProductQuantity = uint(quantity)
	existingProduct.ProductPrice = price
	existingProduct.ProductName = productName

	existingProduct.Update()

	return c.JSON(existingProduct)

}
