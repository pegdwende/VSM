package autorization

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pegdwende/VSM.git/auth/models"
	database "github.com/pegdwende/VSM.git/database"
)

type Config struct {
	Filter       func(c *fiber.Ctx) bool
	Unauthorized fiber.Handler
	Authorized   func(c *fiber.Ctx)
}

var ConfigDefault = Config{
	Filter:       nil,
	Unauthorized: nil,
	Authorized:   nil,
}

func configDefault(config ...Config) Config {

	if len(config) < 1 {
		return ConfigDefault
	}
	cfg := config[0]

	if cfg.Filter != nil {
		cfg.Filter = ConfigDefault.Filter
	}

	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}

	return cfg
}

func New(config Config) fiber.Handler {

	cfg := configDefault(config)

	return func(c *fiber.Ctx) error {

		user := c.Locals("user")
		userInfo, ok := user.(*models.User)

		if !ok {
			cfg.Unauthorized(c)
		}
		var userModel models.User

		database.GetConnection().First(&userModel, "UserName = ?", userInfo.UserName)

		if userModel.ID != 0 {
			return c.Next()
		}

		if config.Filter != nil && config.Filter(c) {

			return c.Next()
		}

		return cfg.Unauthorized(c)
	}
}
