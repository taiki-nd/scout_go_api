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

	app.Get("/api/v1/users", controllers.UsersIndex)
	app.Post("/api/v1/users", controllers.UsersCreate)
	app.Get("/api/v1/users/:id", controllers.UsersShow)
	app.Put("/api/v1/users/:id", controllers.UsersUpdate)
	app.Delete("/api/v1/users/:id", controllers.UsersDelete)

	app.Get("/api/v1/statuses", controllers.StatusesIndex)
	app.Post("/api/v1/statuses", controllers.StatusesCreate)
	app.Get("/api/v1/statuses/:id", controllers.StatusesShow)
	app.Put("/api/v1/statuses/:id", controllers.StatusesUpdate)
	app.Delete("/api/v1/statuses/:id", controllers.StatusesDelete)

	app.Get("/api/v1/prefectures", controllers.PrefecturesIndex)
	app.Post("/api/v1/prefectures", controllers.PrefecturesCreate)
	app.Get("/api/v1/prefectures/:id", controllers.PrefecturesShow)
	app.Put("/api/v1/prefectures/:id", controllers.PrefecturesUpdate)
	app.Delete("/api/v1/prefectures/:id", controllers.PrefecturesDelete)
}
