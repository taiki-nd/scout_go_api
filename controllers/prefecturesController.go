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
 * PrefecturesIndex
 * prefecture一覧を取得
 */
func PrefecturesIndex(c *fiber.Ctx) error {
	log.Println("start to get prefectures")

	var prefectures []*models.Prefecture

	// prefectureesレコードの取得
	err := db.DB.Find(&prefectures).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"prefecture": false,
			"code":       "failed_db_get_prefecturees",
			"message":    fmt.Sprintf("db error: %v", err),
			"data":       prefectures,
		})
	}

	log.Println("success to get prefectures")

	return c.JSON(fiber.Map{
		"prefecture": true,
		"code":       "success_prefecture_index",
		"message":    "",
		"data":       prefectures,
	})
}

/*
 * PrefecturesCreate
 * DBにprefecture情報を登録
 */
func PrefecturesCreate(c *fiber.Ctx) error {
	log.Println("start to create prefecture")

	uuid := c.Query("uuid")

	// admin権限の確認
	err := service.CheckAdmin(uuid)
	if err != nil {
		log.Printf("permission denied: %v", err)
		return c.JSON(fiber.Map{
			"prefecture": false,
			"code":       "permission_error",
			"message":    fmt.Sprintf("permission denied: %v", err),
			"data":       fiber.Map{},
		})
	}

	var prefecture models.Prefecture

	// リクエストボディのパース
	err = c.BodyParser(&prefecture)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"prefecture": false,
			"code":       "failed_parse_prefecture_create",
			"message":    fmt.Sprintf("parse error: %v", err),
			"data":       fiber.Map{},
		})
	}

	// レコード作成
	err = db.DB.Create(&prefecture).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"prefecture": false,
			"code":       "failed_db_prefecture_create",
			"message":    fmt.Sprintf("db error: %v", err),
			"data":       fiber.Map{},
		})
	}

	log.Println("success to create prefecture")

	return c.JSON(fiber.Map{
		"prefecture": true,
		"code":       "success_prefecture_create",
		"message":    "",
		"data":       prefecture,
	})
}

/*
 * PrefecturesShow
 * prefecture詳細情報の取得
 */
func PrefecturesShow(c *fiber.Ctx) error {
	log.Println("start to get prefecture")

	id := c.Params("id")

	// prefecture情報の取得
	prefecture, err := service.GetPrefectureFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"prefecture": false,
			"code":       "failed_db_prefecture_show",
			"message":    fmt.Sprintf("db error: %v", err),
			"data":       fiber.Map{},
		})
	}

	log.Printf("success to get prefectures")

	return c.JSON(fiber.Map{
		"prefecture": true,
		"code":       "success_prefecture_show",
		"message":    "",
		"data":       prefecture,
	})
}

/*
 * PrefecturesUpdate
 * prefecture情報の更新
 */
func PrefecturesUpdate(c *fiber.Ctx) error {
	log.Println("start to Update prefecture")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// prefecture情報の取得
	prefecture, err := service.GetPrefectureFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"prefecture": false,
			"code":       "failed_db_prefecture_show",
			"message":    fmt.Sprintf("db error: %v", err),
			"data":       fiber.Map{},
		})
	}

	// admin権限の確認
	err = service.CheckAdmin(uuid)
	if err != nil {
		log.Printf("permission denied: %v", err)
		return c.JSON(fiber.Map{
			"prefecture": false,
			"code":       "permission_error",
			"message":    fmt.Sprintf("permission denied: %v", err),
			"data":       fiber.Map{},
		})
	}

	// リクエストボディのパース
	err = c.BodyParser(&prefecture)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"prefecture": false,
			"code":       "failed_parse_prefecture_update",
			"message":    fmt.Sprintf("parse error: %v", err),
			"data":       fiber.Map{},
		})
	}

	// prefecture情報の更新
	err = db.DB.Model(&prefecture).Updates(prefecture).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"prefecture": false,
			"code":       "failed_db_prefecture_update",
			"message":    fmt.Sprintf("parse error: %v", err),
			"data":       fiber.Map{},
		})
	}

	log.Println("success to update prefecture")

	return c.JSON(fiber.Map{
		"prefecture": true,
		"code":       "success_prefecture_update",
		"message":    fmt.Sprintf("db error: %v", err),
		"data":       prefecture,
	})
}

/*
 * PrefectureDelete
 * prefecture情報の削除
 */
func PrefecturesDelete(c *fiber.Ctx) error {
	log.Println("start to delete prefecture")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// prefecture情報の取得
	prefecture, err := service.GetPrefectureFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"prefecture": false,
			"code":       "failed_db_prefecture_show",
			"message":    fmt.Sprintf("db error: %v", err),
			"data":       fiber.Map{},
		})
	}

	// admin権限の確認
	err = service.CheckAdmin(uuid)
	if err != nil {
		log.Printf("permission denied: %v", err)
		return c.JSON(fiber.Map{
			"prefecture": false,
			"code":       "permission_error",
			"message":    fmt.Sprintf("permission denied: %v", err),
			"data":       fiber.Map{},
		})
	}

	// prefecture情報の削除
	err = db.DB.Delete(prefecture).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"prefecture": false,
			"code":       "failed_db_prefecture_delete",
			"message":    fmt.Sprintf("db error: %v", err),
			"data":       fiber.Map{},
		})
	}

	log.Println("success to delete prefecture")

	return c.JSON(fiber.Map{
		"prefecture": true,
		"code":       "success_prefecture_delete",
		"message":    "",
		"data":       fiber.Map{},
	})
}
