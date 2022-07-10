package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pegdwende/VSM.git/auth/models"
	"github.com/pegdwende/VSM.git/database"
	"github.com/pegdwende/VSM.git/env"
	"golang.org/x/crypto/bcrypt"
)

//TODO create a sperate route to register admin users
func RegisterUser(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	adminUser, _ := strconv.ParseBool(data["admin_user"])

	roleId, _ := strconv.Atoi(data["role_id"])

	user := models.User{
		UserName:   data["username"],
		Email:      data["email"],
		Phone:      data["phone_number"],
		ClientCode: data["client_code"],
		RoleId:     uint(roleId),
		AdminUser:  adminUser,
		Password:   string(password),
		UserStatus: models.USER_ACTIVE_STATUS,
	}

	database.GetConnection().Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.GetConnection().Where("user_name = ? AND client_code = ? ", data["user_name"], data["client_code"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		c.JSON(fiber.Map{"message": "user not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusNotFound)
		c.JSON(fiber.Map{"message": "User not found"})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_nane":   user.UserName,
		"admin":       user.AdminUser,
		"role_id":     user.RoleId,
		"client_code": user.ClientCode,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(env.GetRequiredEnvVariable("JWT_SECRET")))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
