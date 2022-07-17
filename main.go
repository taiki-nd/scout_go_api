package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/taiki-nd/scout_go_api/config"
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/utils"
)

func main() {
	// log出力の有効化
	utils.Logging(config.Config.LogFile)

	// db接続
	db.ConnectToDb()

	// fiber(フレームワーク接続)
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Println("starting server at port:8000")

	app.Listen(":8000")
}
