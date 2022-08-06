package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
	"github.com/taiki-nd/scout_go_api/service"
	"gorm.io/gorm"
)

/*
 * ResumesIndex
 * resume一覧を取得
 */
func ResumesIndex(c *fiber.Ctx) error {
	log.Println("start to get resumes")

	var resumes []*models.Resume

	// resumesレコードの取得
	err := db.DB.Preload("Projects").Find(&resumes).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_get_resumees",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    resumes,
		})
	}

	log.Println("success to get resumes")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_resume_index",
		"message": "",
		"data":    resumes,
	})
}

/*
 * ResumesCreate
 * DBにresume情報を登録
 */
func ResumesCreate(c *fiber.Ctx) error {
	log.Println("start to create resume")

	uuid := c.Query("uuid")
	if uuid == "" {
		log.Println("not_signin")
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "not_signin",
			"message": "please signin",
			"data":    fiber.Map{},
		})
	}

	var resume models.Resume

	// リクエストボディのパース
	err := c.BodyParser(&resume)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_parse_resume_create",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// レコード作成
	err = db.DB.Create(&resume).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_resume_create",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to create resume")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_resume_create",
		"message": "",
		"data":    resume,
	})
}

/*
 * ResumesShow
 * resume詳細情報の取得
 */
func ResumesShow(c *fiber.Ctx) error {
	log.Println("start to get resume")

	id := c.Params("id")

	// resume情報の取得
	resume, err := service.GetResumeFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_resume_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Printf("success to get resumes")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_resume_show",
		"message": "",
		"data":    resume,
	})
}

/*
 * ResumesUpdate
 * resume情報の更新
 */
func ResumesUpdate(c *fiber.Ctx) error {
	log.Println("start to Update resume")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// resume情報の取得
	resume, err := service.GetResumeFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_resume_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// userステータスの確認
	err = service.CheckUserStatus(uuid, resume.UserId)
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
	err = c.BodyParser(&resume)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_parse_resume_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// resume情報の更新
	err = db.DB.Model(&resume).Updates(resume).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_resume_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to update resume")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_resume_update",
		"message": fmt.Sprintf("db error: %v", err),
		"data":    resume,
	})
}

/*
 * ResumesDelete
 * resume情報の削除
 */
func ResumesDelete(c *fiber.Ctx) error {
	log.Println("start to delete resume")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// resume情報の取得
	resume, err := service.GetResumeFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_resume_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// userステータスの確認
	err = service.CheckUserStatus(uuid, resume.UserId)
	if err != nil {
		log.Printf("user stratus error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "user_status_error",
			"message": fmt.Sprintf("user stratus error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// transaction開始
	errTransaction := db.DB.Transaction(func(tx *gorm.DB) error {
		errProject := tx.Table("projects").Where("resume_id = ?", resume.Id).Delete("").Error
		if errProject != nil {
			log.Printf("db error: %v", errProject)
			return fmt.Errorf("db error: %v", errProject)
		}
		// resume情報の削除
		err = tx.Delete(resume).Error
		if err != nil {
			log.Printf("db error: %v", err)
			return fmt.Errorf("db error: %v", err)
		}
		return nil
	})
	if errTransaction != nil {
		log.Println(errTransaction)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_resume_delete",
			"message": fmt.Sprintf("db error: %v", errTransaction),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to delete resume")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_resume_delete",
		"message": "",
		"data":    fiber.Map{},
	})
}
