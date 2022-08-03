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
 * StatusesIndex
 * status一覧を取得
 */
func StatusesIndex(c *fiber.Ctx) error {
	log.Println("start to get statuses")

	var statuses []*models.Status

	// statusesレコードの取得
	err := db.DB.Find(&statuses).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_get_statuses",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    statuses,
		})
	}

	log.Println("success to get statuses")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_status_index",
		"message": "",
		"data":    statuses,
	})
}

/*
 * StatusesCreate
 * DBにstatus情報を登録
 */
func StatusesCreate(c *fiber.Ctx) error {
	log.Println("start to create status")

	uuid := c.Params("uuid")

	// admin権限の確認
	err := service.CheckAdmin(uuid)
	if err != nil {
		log.Printf("permission denied: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "permission_error",
			"message": fmt.Sprintf("permission denied: %v", err),
			"data":    fiber.Map{},
		})
	}

	var status models.Status

	// リクエストボディのパース
	err = c.BodyParser(&status)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_parse_status_create",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// レコード作成
	err = db.DB.Create(&status).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_status_create",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to create status")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_status_create",
		"message": "",
		"data":    status,
	})
}

/*
 * StatusesShow
 * status詳細情報の取得
 */
func StatusesShow(c *fiber.Ctx) error {
	log.Println("start to get status")

	id := c.Params("id")

	// status情報の取得
	status, err := service.GetStatusFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_status_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Printf("success to get statuses")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_status_show",
		"message": "",
		"data":    status,
	})
}

/*
 * StatusesUpdate
 * status情報の更新
 */
func StatusesUpdate(c *fiber.Ctx) error {
	log.Println("start to Update status")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// status情報の取得
	status, err := service.GetStatusFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_status_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// admin権限の確認
	err = service.CheckAdmin(uuid)
	if err != nil {
		log.Printf("permission denied: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "permission_error",
			"message": fmt.Sprintf("permission denied: %v", err),
			"data":    fiber.Map{},
		})
	}

	// リクエストボディのパース
	err = c.BodyParser(&status)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_parse_status_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// status情報の更新
	err = db.DB.Model(&status).Updates(status).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_status_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to update status")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_status_update",
		"message": fmt.Sprintf("db error: %v", err),
		"data":    status,
	})
}

/*
 * StatusesDelete
 * status情報の削除
 */
func StatusesDelete(c *fiber.Ctx) error {
	log.Println("start to delete status")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// status情報の取得
	status, err := service.GetStatusFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_status_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// admin権限の確認
	err = service.CheckAdmin(uuid)
	if err != nil {
		log.Printf("permission denied: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "permission_error",
			"message": fmt.Sprintf("permission denied: %v", err),
			"data":    fiber.Map{},
		})
	}

	// status情報の削除
	err = db.DB.Delete(status).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_status_delete",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to delete status")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_status_delete",
		"message": "",
		"data":    fiber.Map{},
	})
}
