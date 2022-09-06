package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pegdwende/VSM.git/auth/models"
	AuthModels "github.com/pegdwende/VSM.git/auth/models"
	database "github.com/pegdwende/VSM.git/database"
)

func RegisterClient(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var existingClient models.Client

	database.GetConnection().Where("client_code = ?", data["code"]).First(&existingClient)

	if existingClient.ID != 0 {
		return c.JSON(fiber.Map{"message": "client code exist"})
	}

	client := AuthModels.Client{
		ClientCode:   data["code"],
		Name:         data["name"],
		Address:      data["address"],
		Email:        data["email"],
		ClientStatus: AuthModels.P,
		BusinessType: AuthModels.RETAIL,
	}

	client.Create()

	return c.JSON(client)

}
