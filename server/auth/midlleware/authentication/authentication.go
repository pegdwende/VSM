package authentication

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pegdwende/VSM.git/auth/models"
	database "github.com/pegdwende/VSM.git/database"
	"github.com/pegdwende/VSM.git/env"
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

		jwtCookie := c.Cookies("jwt")

		token, err := jwt.ParseWithClaims(jwtCookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(env.GetRequiredEnvVariable("JWT_SECRET")), nil
		})

		if err != nil {
			cfg.Unauthorized(c)
		}

		claims := token.Claims.(*jwt.StandardClaims)
		var user models.User

		database.GetConnection().Where("id = ?", claims.Issuer).First(&user)
		if user.ID != 0 {

			c.Locals("user", user)
			return c.Next()
		}

		// if !ok {
		// 	cfg.Unauthorized(c)
		// }
		// var userModel models.User

		// database.GetConnection().First(&userModel, "UserName = ?", userInfo.UserName)

		if config.Filter != nil && config.Filter(c) {

			return c.Next()
		}

		return cfg.Unauthorized(c)
	}
}
