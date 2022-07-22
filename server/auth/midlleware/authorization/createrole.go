package autorization

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pegdwende/VSM.git/auth/models"
	"github.com/pegdwende/VSM.git/database"
)

func CreateRole(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["user_nane"].(string)

	var existingUser models.User

	database.GetConnection().Where("user_name = ?", name).Preload("Role.Permissions").First(&existingUser)

	userRole := existingUser.Role

	fmt.Println(userRole.Permissions)

	rolePermissions := userRole.Permissions

	for _, element := range rolePermissions {

		fmt.Println(element)
		if element.PermissionKey == models.CREATE_ROLE {

			fmt.Println("user is able to create roles")
		}
	}
	return c.Next()

}
