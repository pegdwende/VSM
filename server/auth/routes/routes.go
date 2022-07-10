package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pegdwende/VSM.git/auth/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register-clients", controllers.RegisterClient)
	app.Post("/api/create-role", controllers.CreateRole)
	app.Post("/api/assigned-role-permissions", controllers.AssignRolePermission)
	app.Get("/api/role/:roleId", controllers.RoleAndPermissions)
	app.Post("api/register-user", controllers.RegisterUser)
	app.Post("api/login", controllers.Login)
	app.Post("api/logout", controllers.Logout)
}
