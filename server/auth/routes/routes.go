package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pegdwende/VSM.git/auth/controllers"
	autorization "github.com/pegdwende/VSM.git/auth/midlleware/authorization"
)

func testToken(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["user_nane"].(string)

	fmt.Println(name)
	return c.Next()
}
func Setup(app *fiber.App) {
	// app.Post("/api/register-clients", controllers.RegisterClient)
	app.Post("/api/create-role", autorization.CreateRole, controllers.CreateRole)
	app.Post("/api/assigned-role-permissions", controllers.AssignRolePermission)
	app.Get("/api/role/:roleId", controllers.RoleAndPermissions)
	// app.Post("api/register-user", controllers.RegisterUser)
	// app.Post("api/login", controllers.Login)
	app.Post("api/logout", controllers.Logout)
}

func test(c *fiber.Ctx) error {
	fmt.Println("middleware works here man")

	return c.Next()
}

func SetupPublicRoutes(app *fiber.App) {
	app.Post("/api/register-clients", test, controllers.RegisterClient)
	app.Post("api/register-user", controllers.RegisterUser)
	app.Post("api/login", controllers.Login)
}
