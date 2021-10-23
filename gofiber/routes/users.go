package routes

import (
	"gofiber/config"
	"gofiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func Users(app *fiber.App) {
	db := config.InitDB()
	usersGroup := app.Group("/api/v1/users")
	usersController := controllers.Users{DB : db}
	{
		usersGroup.Get("/", usersController.FindAll)
		usersGroup.Get("/:id", usersController.FindOne)
		//usersGroup.Get("/limit/test/", usersController.FindLimit)
		usersGroup.Patch("/:id", usersController.Update)
		usersGroup.Delete("/:id", usersController.Delete)
		usersGroup.Post("", usersController.Create)
	}

	v1 := app.Group("/api/v1")
	authGroup := v1.Group("/auth")
	authController := controllers.Auth{DB: db}
	{
		authGroup.Post("/signup", authController.Signup)
	}
}