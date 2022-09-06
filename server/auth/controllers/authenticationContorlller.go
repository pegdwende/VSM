package controllers

import (
	"fmt"
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

	if len(data["password"]) < 1 || len(data["username"]) < 1 {
		return c.JSON(fiber.Map{"message": "User and password must be provided "})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	adminUser, _ := strconv.ParseBool(data["admin_user"])

	roleId, _ := strconv.Atoi(data["role_id"])
	clientId, _ := strconv.Atoi(data["client_id"])

	var existingUserModel models.User

	database.GetConnection().Where("user_name = ?", data["username"]).First(&existingUserModel)

	if existingUserModel.ID != 0 {
		fmt.Println(existingUserModel)
		return c.JSON(fiber.Map{"message": "user already exist."})
	}

	var role models.Role

	database.GetConnection().Where("id=?", roleId).First(&role)

	if role.ID == 0 {
		return c.JSON(fiber.Map{"message": "Provided role does not exist."})
	}

	user := models.User{
		UserName:   data["username"],
		Email:      data["email"],
		Phone:      data["phone_number"],
		ClientID:   uint(clientId),
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

	database.GetConnection().Where("user_name = ?", data["user_name"]).Preload("Client", "client_code = ?", data["client_code"]).First(&user)

	if user.ID == 0 || user.Client.ID == 0 {
		c.Status(fiber.StatusNotFound)
		c.JSON(fiber.Map{"message": "user not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusNotFound)
		c.JSON(fiber.Map{"message": "User not found"})
	}

	fmt.Println(user)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_name": user.UserName,
		"admin":     user.AdminUser,
		"role_id":   user.RoleId,
		"client_id": user.ClientID,
		"user_id":   user.ID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	fmt.Println(claims)

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
