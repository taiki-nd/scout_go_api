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
 * SchoolsIndex
 * school一覧を取得
 */
func SchoolsIndex(c *fiber.Ctx) error {
	log.Println("start to get schools")

	var schools []*models.School

	// schoolsレコードの取得
	err := db.DB.Find(&schools).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"school":  false,
			"code":    "failed_db_get_schooles",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    schools,
		})
	}

	log.Println("success to get schools")

	return c.JSON(fiber.Map{
		"school":  true,
		"code":    "success_school_index",
		"message": "",
		"data":    schools,
	})
}

/*
 * SchoolsCreate
 * DBにschool情報を登録
 */
func SchoolsCreate(c *fiber.Ctx) error {
	log.Println("start to create school")

	uuid := c.Query("uuid")
	if uuid == "" {
		log.Println("not_signin")
		return c.JSON(fiber.Map{
			"school":  false,
			"code":    "not_signin",
			"message": "please signin",
			"data":    fiber.Map{},
		})
	}

	var school models.School

	// リクエストボディのパース
	err := c.BodyParser(&school)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"school":  false,
			"code":    "failed_parse_school_create",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// レコード作成
	err = db.DB.Create(&school).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"school":  false,
			"code":    "failed_db_school_create",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to create school")

	return c.JSON(fiber.Map{
		"school":  true,
		"code":    "success_school_create",
		"message": "",
		"data":    school,
	})
}

/*
 * SchoolsShow
 * school詳細情報の取得
 */
func SchoolsShow(c *fiber.Ctx) error {
	log.Println("start to get school")

	id := c.Params("id")

	// school情報の取得
	school, err := service.GetSchoolFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"school":  false,
			"code":    "failed_db_school_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Printf("success to get schools")

	return c.JSON(fiber.Map{
		"school":  true,
		"code":    "success_school_show",
		"message": "",
		"data":    school,
	})
}

/*
 * SchoolsUpdate
 * school情報の更新
 */
func SchoolsUpdate(c *fiber.Ctx) error {
	log.Println("start to Update school")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// school情報の取得
	school, err := service.GetSchoolFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"school":  false,
			"code":    "failed_db_school_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// userステータスの確認
	err = service.CheckUserStatus(uuid, school.UserId)
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
	err = c.BodyParser(&school)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"school":  false,
			"code":    "failed_parse_school_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// school情報の更新
	err = db.DB.Model(&school).Updates(school).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"school":  false,
			"code":    "failed_db_school_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to update school")

	return c.JSON(fiber.Map{
		"school":  true,
		"code":    "success_school_update",
		"message": fmt.Sprintf("db error: %v", err),
		"data":    school,
	})
}

/*
 * SchoolsDelete
 * school情報の削除
 */
func SchoolsDelete(c *fiber.Ctx) error {
	log.Println("start to delete school")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// school情報の取得
	school, err := service.GetSchoolFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"school":  false,
			"code":    "failed_db_school_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// userステータスの確認
	err = service.CheckUserStatus(uuid, school.UserId)
	if err != nil {
		log.Printf("user stratus error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "user_status_error",
			"message": fmt.Sprintf("user stratus error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// school情報の削除
	err = db.DB.Delete(school).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"school":  false,
			"code":    "failed_db_school_delete",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to delete school")

	return c.JSON(fiber.Map{
		"school":  true,
		"code":    "success_school_delete",
		"message": "",
		"data":    fiber.Map{},
	})
}
