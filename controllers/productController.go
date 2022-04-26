package controllers

import (
	"fmt"
	"strconv"

	"../database"

	"../models"

	"github.com/gofiber/fiber"
)

func CreateProduct(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {

		return err

	}

	// categories := database.DB.Find(&models.Category);
	category, err := strconv.ParseUint(data["category"], 0, 90)
	status, err := strconv.ParseUint(data["status"], 0, 90)
	product := models.Product{
		Name:           data["name"],
		Price:          data["price"],
		Quantity:       data["quantity"],
		Description:    data["description"],
		CategoryRefer:  uint(category),
		ProductStRefer: uint(status),
	}

	if err == nil {
		fmt.Print(category)
	}

	database.DB.Create(&product)

	return c.JSON(product)
}

func GetProducts(c *fiber.Ctx) error {
	var product []models.Product

	database.DB.Find(&product)

	return c.JSON(&product)
}
