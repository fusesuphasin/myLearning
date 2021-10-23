package main

import (
	"gofiber/config"
	"gofiber/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}

	config.InitDB()

	app := fiber.New()

	app.Static("/uploads", "./uploads") 

	uploadDirs := [...]string{"users"}

	for _, dir := range uploadDirs {
		os.MkdirAll("uploads/"+dir, 0755 /*Unix permission calculate*/)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Wellcome!")
	})

	
	routes.Getvalue(app)
	routes.Users(app)

	app.Listen(":3000")
}
