package controllers

import (
	"../database"

	"../models"

	"github.com/gofiber/fiber"
)

func GetAllCategory(c *fiber.Ctx) error {
	var categories []models.Category

	database.DB.Find(&categories)

	return c.JSON(categories)
}

func CreateCategory(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusUnauthorized)

		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	category := models.Category{
		Name: data["name"],
	}

	createCategory := database.DB.Create(&category)

	return c.JSON(createCategory)
}
