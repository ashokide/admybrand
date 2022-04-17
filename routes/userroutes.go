package routes

import (
	"admybrand/controllers"

	"github.com/gofiber/fiber"
)

func UserRoutes(app *fiber.App) {
	app.Get("/", controllers.HelloWorld)
	app.Get("/api/users", controllers.GetAllUsers)
	app.Get("/api/user/:id", controllers.GetAUser)
	app.Post("/api/user/insert", controllers.InsertAUser)
	app.Put("/api/user/update/:id", controllers.UpdateAUser)
	app.Delete("/api/user/delete/:id", controllers.DeleteAUser)
}
