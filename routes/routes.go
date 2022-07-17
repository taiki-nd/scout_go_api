package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/taiki-nd/scout_go_api/controllers"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.UsersIndex)
}
