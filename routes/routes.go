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

	app.Post("api/v1/images", controllers.ImageUpload)

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

	app.Get("/api/v1/licenses", controllers.LicensesIndex)
	app.Post("/api/v1/licenses", controllers.LicensesCreate)
	app.Get("/api/v1/licenses/:id", controllers.LicensesShow)
	app.Put("/api/v1/licenses/:id", controllers.LicensesUpdate)
	app.Delete("/api/v1/licenses/:id", controllers.LicensesDelete)

	app.Get("/api/v1/schools", controllers.SchoolsIndex)
	app.Post("/api/v1/schools", controllers.SchoolsCreate)
	app.Get("/api/v1/schools/:id", controllers.SchoolsShow)
	app.Put("/api/v1/schools/:id", controllers.SchoolsUpdate)
	app.Delete("/api/v1/schools/:id", controllers.SchoolsDelete)

	app.Get("/api/v1/activities", controllers.ActivitiesIndex)
	app.Post("/api/v1/activities", controllers.ActivitiesCreate)
	app.Get("/api/v1/activities/:id", controllers.ActivitiesShow)
	app.Put("/api/v1/activities/:id", controllers.ActivitiesUpdate)
	app.Delete("/api/v1/activities/:id", controllers.ActivitiesDelete)

	app.Get("/api/v1/works", controllers.WorksIndex)
	app.Post("/api/v1/works", controllers.WorksCreate)
	app.Get("/api/v1/works/:id", controllers.WorksShow)
	app.Put("/api/v1/works/:id", controllers.WorksUpdate)
	app.Delete("/api/v1/works/:id", controllers.WorksDelete)
}
