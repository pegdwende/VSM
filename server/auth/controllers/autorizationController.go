package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	AuthModels "github.com/pegdwende/VSM.git/auth/models"
	database "github.com/pegdwende/VSM.git/database"
	"gorm.io/gorm/clause"
)

func CreateRole(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	clientId, _ := strconv.Atoi(data["client_id"])
	role := AuthModels.Role{
		RoleDescription: data["role_description"],
		ClientID:        uint(clientId),
	}

	database.GetConnection().Create(&role)

	return c.JSON(role)

}

func AssignRolePermission(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	roleId, err := strconv.Atoi(data["role_id"])

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.SendString("Invalid role id provided")
	}

	clientId, _ := strconv.Atoi(data["client_id"])
	role := AuthModels.Permission{
		ClientID:              uint(clientId),
		PermissionKey:         data["permission_key"],
		RoleID:                uint(roleId),
		PermissionDescription: data["permission_desc"],
	}

	database.GetConnection().Create(&role)

	return c.JSON(role)
}

func RoleAndPermissions(c *fiber.Ctx) error {
	roleIdString := c.Params("roleId")

	roleId, err := strconv.Atoi(roleIdString)

	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		c.SendString("Invalid role id provided")
	}

	var role AuthModels.Role

	database.GetConnection().Preload(clause.Associations).First(&role, roleId)

	return c.JSON(role.Permissions)
}
