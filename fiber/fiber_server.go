package fiber

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func Start(path, port string) {
	app := fiber.New()

	app.Use(middleware.Logger())

	app.Static("/s/*", path)

	app.Get("/t/:fileName", func(c *fiber.Ctx) {
		fileName := c.Params("fileName")
		fullPath := path + "/t/" + fileName
		fmt.Printf("fullPath: %s\n", fullPath)
		b, _ := ioutil.ReadFile(fullPath)
		utf8Str, _ := ConvertGBK2Str(string(b))
		s := strings.Replace(utf8Str, "\r\n\r\n", "\r\n", -1)
		s = strings.Replace(s, "\r\n\r\n", "\r\n", -1)
		fmt.Printf("I'm here: %s\n", s)
		c.Send(s)
	})

	// // Respond with Hello, World! on the homepage:
	// app.Get("/", func(c *fiber.Ctx) {
	// 	c.Send("Hello, World!")
	// })

	// // Parameter
	// // http://localhost:8080/hello%20world
	// app.Post("/:value", func(c *fiber.Ctx) {
	// 	c.Send("Post request with value: " + c.Params("value"))
	// 	// => Post request with value: hello world
	// })

	// // Optional parameter
	// // http://localhost:8080/hello%20world
	// app.Get("/:value?", func(c *fiber.Ctx) {
	// 	if c.Params("value") != "" {
	// 		c.Send("Get request with value: " + c.Params("value"))
	// 		return // => Get request with value: hello world
	// 	}
	// 	c.Send("Get request without value")
	// })

	// // Wildcard
	// // http://localhost:8080/api/user/john
	// app.Get("/api/*", func(c *fiber.Ctx) {
	// 	c.Send("API path with wildcard: " + c.Params("*"))
	// 	// => API path with wildcard: user/john
	// })

	app.Listen(port)
}

func ConvertGBK2Str(gbkStr string) (string, error) {
	return simplifiedchinese.GBK.NewDecoder().String(gbkStr)
}
