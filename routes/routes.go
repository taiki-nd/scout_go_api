package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/taiki-nd/scout_go_api/controllers"
)

func Routes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders: "Content-Type, Content-Length, Authorization, Uid",
		AllowOrigins: "*",
	}))
	app.Get("/", controllers.UsersIndex)
	app.Post("/api/v1/users", controllers.UsersCreate)
}
