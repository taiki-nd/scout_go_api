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
 * WorksIndex
 * work一覧を取得
 */
func WorksIndex(c *fiber.Ctx) error {
	log.Println("start to get works")

	var works []*models.Work

	// worksレコードの取得
	err := db.DB.Preload("Projects").Find(&works).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"work":    false,
			"code":    "failed_db_get_workes",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    works,
		})
	}

	log.Println("success to get works")

	return c.JSON(fiber.Map{
		"work":    true,
		"code":    "success_work_index",
		"message": "",
		"data":    works,
	})
}

/*
 * WorksCreate
 * DBにwork情報を登録
 */
func WorksCreate(c *fiber.Ctx) error {
	log.Println("start to create work")

	uuid := c.Query("uuid")
	if uuid == "" {
		log.Println("not_signin")
		return c.JSON(fiber.Map{
			"work":    false,
			"code":    "not_signin",
			"message": "please signin",
			"data":    fiber.Map{},
		})
	}

	var work models.Work

	// リクエストボディのパース
	err := c.BodyParser(&work)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"work":    false,
			"code":    "failed_parse_work_create",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// レコード作成
	err = db.DB.Create(&work).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"work":    false,
			"code":    "failed_db_work_create",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to create work")

	return c.JSON(fiber.Map{
		"work":    true,
		"code":    "success_work_create",
		"message": "",
		"data":    work,
	})
}

/*
 * WorksShow
 * work詳細情報の取得
 */
func WorksShow(c *fiber.Ctx) error {
	log.Println("start to get work")

	id := c.Params("id")

	// work情報の取得
	work, err := service.GetWorkFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"work":    false,
			"code":    "failed_db_work_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Printf("success to get works")

	return c.JSON(fiber.Map{
		"work":    true,
		"code":    "success_work_show",
		"message": "",
		"data":    work,
	})
}

/*
 * WorksUpdate
 * work情報の更新
 */
func WorksUpdate(c *fiber.Ctx) error {
	log.Println("start to Update work")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// work情報の取得
	work, err := service.GetWorkFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"work":    false,
			"code":    "failed_db_work_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// userステータスの確認
	err = service.CheckUserStatus(uuid, work.UserId)
	if err != nil {
		log.Printf("user stratus error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "user_status_error",
			"message": fmt.Sprintf("user stratus error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// リクエストボディのパース
	err = c.BodyParser(&work)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"work":    false,
			"code":    "failed_parse_work_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// work情報の更新
	err = db.DB.Model(&work).Updates(work).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"work":    false,
			"code":    "failed_db_work_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to update work")

	return c.JSON(fiber.Map{
		"work":    true,
		"code":    "success_work_update",
		"message": fmt.Sprintf("db error: %v", err),
		"data":    work,
	})
}

/*
 * WorksDelete
 * work情報の削除
 */
func WorksDelete(c *fiber.Ctx) error {
	log.Println("start to delete work")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// work情報の取得
	work, err := service.GetWorkFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"work":    false,
			"code":    "failed_db_work_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// userステータスの確認
	err = service.CheckUserStatus(uuid, work.UserId)
	if err != nil {
		log.Printf("user stratus error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "user_status_error",
			"message": fmt.Sprintf("user stratus error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// work情報の削除
	err = db.DB.Delete(work).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"work":    false,
			"code":    "failed_db_work_delete",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to delete work")

	return c.JSON(fiber.Map{
		"work":    true,
		"code":    "success_work_delete",
		"message": "",
		"data":    fiber.Map{},
	})
}
