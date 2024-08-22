package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"test.com/firstgoproject/internal/handlers"
)

func healthchek(c *fiber.Ctx) error {
	return c.SendString("all wright")
}

func main() {
	app := fiber.New()

	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("hellow form midl ware")
		return c.Next()
	})

	app.Get("healthchek", healthchek)
	app.Post("/api/products", handlers.CreatProduct)
	app.Post("/api/user",handlers.CreatUsers)
	app.Get("/api/products", handlers.GetAllProducts)
	fmt.Println("console.log main fill is working")


	log.Fatal(app.Listen(":3000"))

}
