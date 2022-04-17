package main

import (
	"admybrand/configs"
	"admybrand/routes"
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	fmt.Println("Started...!")
	app := fiber.New()

	routes.UserRoutes(app)
	app.Listen(configs.GetPort())
}
