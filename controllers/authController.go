package controllers

import (
	"../database"

	"../models"

	"github.com/dgrijalva/jwt-go"

	"github.com/gofiber/fiber"

	"golang.org/x/crypto/bcrypt"

	"strconv"

	"time"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {

		return err

	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{}
	database.DB.Create(&user)

	account := models.Account{

		Name: data["name"],

		Email: data["email"],

		Password: password,

		UserID: user.ID,
	}

	database.DB.Create(&account)
	return c.JSON(account)

}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	var category models.Category

	categories := database.DB.Find(&category)

	println(categories)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var account models.Account

	database.DB.Where("email = ?", data["email"]).First(&account)

	if account.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "account not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(account.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(account.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
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

func Account(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)

		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var account models.Account

	database.DB.Where("id = ?", claims.Issuer).First(&account)

	return c.JSON(account)
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
