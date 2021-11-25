package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", "")
	})
	log.Fatal(app.Listen(":3000"))
}
