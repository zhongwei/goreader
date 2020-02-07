package fiber

import "github.com/gofiber/fiber"

func Start() {
	app := fiber.New()

	// Respond with Hello, World! on the homepage:
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	// Parameter
	// http://localhost:8080/hello%20world
	app.Post("/:value", func(c *fiber.Ctx) {
		c.Send("Post request with value: " + c.Params("value"))
		// => Post request with value: hello world
	})

	// Optional parameter
	// http://localhost:8080/hello%20world
	app.Get("/:value?", func(c *fiber.Ctx) {
		if c.Params("value") != "" {
			c.Send("Get request with value: " + c.Params("value"))
			return // => Get request with value: hello world
		}
		c.Send("Get request without value")
	})

	// Wildcard
	// http://localhost:8080/api/user/john
	app.Get("/api/*", func(c *fiber.Ctx) {
		c.Send("API path with wildcard: " + c.Params("*"))
		// => API path with wildcard: user/john
	})

	app.Listen(8080)
}
