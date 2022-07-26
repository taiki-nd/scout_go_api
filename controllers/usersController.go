package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
	"github.com/taiki-nd/scout_go_api/service"
)

/*
 * UsersIndex
 * user一覧を取得
 */
func UsersIndex(c *fiber.Ctx) error {
	log.Println("start to get users")

	var users []*models.User

	// usersレコードの取得
	err := db.DB.Find(&users).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_get_users",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    users,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_user_index",
		"message": "",
		"data":    users,
	})
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
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_parse_user_create",
			"message": fmt.Sprintf("parse error: %v", err),
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

/*
 * UsersShow
 * user詳細情報の取得
 */
func UsersShow(c *fiber.Ctx) error {
	log.Println("start to get user")

	// user情報の取得
	user, err := service.GetUserFromId(c)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_user_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_user_show",
		"message": "",
		"data":    user,
	})
}

/*
 * UsersUpdate
 * user情報の更新
 */
func UsersUpdate(c *fiber.Ctx) error {
	log.Println("start to Update user")

	// user情報の取得
	user, err := service.GetUserFromId(c)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_user_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// リクエストボディのパース
	err = c.BodyParser(&user)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_parse_user_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// user情報の更新
	err = db.DB.Model(&user).Updates(user).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_user_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	return c.JSON(fiber.Map{
		"status":  false,
		"code":    "failed_db_user_update",
		"message": fmt.Sprintf("db error: %v", err),
		"data":    user,
	})
}
