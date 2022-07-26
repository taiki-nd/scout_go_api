package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
)

func UsersIndex(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

/*
 * UsersCreate
 * DBにuser情報を登録
 */
func UsersCreate(c *fiber.Ctx) error {
	log.Println("start to create user")

	var user models.User

	// リクエストボディのパース
	err := c.BodyParser(&user)
	if err != nil {
		log.Printf("POST method error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_parse_user_create",
			"message": fmt.Sprintf("POST method error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// レコード作成
	err = db.DB.Create(&user).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_user_create",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Printf("success to create user: %v", user.Nickname)

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_user_create",
		"message": "",
		"data":    user,
	})
}
