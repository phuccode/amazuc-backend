package routes

import (
	"../controllers"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {

	//auth
	auth := app.Group("/api")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	auth.Get("/user", controllers.Account)
	auth.Post("/logout", controllers.Logout)

	//product
	product := app.Group("/product")
	product.Get("/product", controllers.GetProducts)
	//category
	category := app.Group("/category")
	category.Get("/list", controllers.GetAllCategory)
	category.Post("/create", controllers.CreateCategory)
}
