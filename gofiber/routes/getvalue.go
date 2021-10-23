package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Getvalue(app *fiber.App){

	getValue := app.Group("/api/v1/getValue")

	// GET http://localhost:3000/john
	getValue.Get("/:name?", func(c *fiber.Ctx) error {
		if name := c.Params("name"); name != "" {
			return c.SendString("Hello " + name)
			// => Hello john
		}
		return c.SendString("Who is it?")
	})

	// GET http://localhost:3000/api/user/john

	getValue.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
		// => API path: user/john
	})

	getValue.Get("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil{
			return fiber.ErrBadRequest
		}
		return c.SendString("ID: " + strconv.Itoa(id))
	})

	getValue.Get("/query/name", func(c *fiber.Ctx) error {
		name := c.Query("name")
		surname := c.Query("surname")
		return c.SendString("name: " + name + " surname: " + surname)
	})
}