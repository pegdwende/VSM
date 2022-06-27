package controllers

import (
	"github.com/gofiber/fiber/v2"
	AuthModels "github.com/pegdwende/VSM.git/auth/models"
	database "github.com/pegdwende/VSM.git/database"
)

func RegisterClient(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	client := AuthModels.Client{
		Code:         data["code"],
		Name:         data["name"],
		Address:      data["address"],
		Email:        data["email"],
		ClientStatus: AuthModels.P,
		BusinessType: AuthModels.RETAIL,
	}

	database.GetConnection().Create(&client)

	return c.JSON(client)

}
