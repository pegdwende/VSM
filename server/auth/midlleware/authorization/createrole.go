package autorization

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pegdwende/VSM.git/auth/models"
	"github.com/pegdwende/VSM.git/database"
	"gorm.io/gorm/clause"
)

func CreateRole(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["user_nane"].(string)

	var existingUser models.User

	database.GetConnection().Where("user_name = ?", name).Preload(clause.Associations).First(&existingUser)

	fmt.Println(existingUser.Role)
	fmt.Println(existingUser)
	return c.Next()

}
