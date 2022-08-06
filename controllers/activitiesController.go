package controllers

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/taiki-nd/scout_go_api/config"
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
	"github.com/taiki-nd/scout_go_api/service"
)

/*
 * ActivitiesIndex
 * activity一覧を取得
 */
func ActivitiesIndex(c *fiber.Ctx) error {
	log.Println("start to get activities")

	var activities []*models.Activity

	// activityesレコードの取得
	err := db.DB.Find(&activities).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_get_activityes",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    activities,
		})
	}

	log.Println("success to get activities")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_activity_index",
		"message": "",
		"data":    activities,
	})
}

/*
 * ActivitiesCreate
 * DBにactivity情報を登録
 */
func ActivitiesCreate(c *fiber.Ctx) error {
	log.Println("start to create activity")

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

	var activity models.Activity

	// リクエストボディのパース
	err := c.BodyParser(&activity)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_parse_activity_create",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// レコード作成
	err = db.DB.Create(&activity).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_activity_create",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to create activity")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_activity_create",
		"message": "",
		"data":    activity,
	})
}

/*
 * ActivitiesShow
 * activity詳細情報の取得
 */
func ActivitiesShow(c *fiber.Ctx) error {
	log.Println("start to get activity")

	id := c.Params("id")

	// activity情報の取得
	activity, err := service.GetActivityFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_activity_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Printf("success to get activities")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_activity_show",
		"message": "",
		"data":    activity,
	})
}

/*
 * ActivitiesUpdate
 * activity情報の更新
 */
func ActivitiesUpdate(c *fiber.Ctx) error {
	log.Println("start to Update activity")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// activity情報の取得
	activity, err := service.GetActivityFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_activity_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// userステータスの確認
	err = service.CheckUserStatus(uuid, activity.UserId)
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
	err = c.BodyParser(&activity)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_parse_activity_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// activity情報の更新
	err = db.DB.Model(&activity).Updates(activity).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_activity_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to update activity")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_activity_update",
		"message": fmt.Sprintf("db error: %v", err),
		"data":    activity,
	})
}

/*
 * ActivitiesDelete
 * activity情報の削除
 */
func ActivitiesDelete(c *fiber.Ctx) error {
	log.Println("start to delete activity")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// activity情報の取得
	activity, err := service.GetActivityFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_activity_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// userステータスの確認
	err = service.CheckUserStatus(uuid, activity.UserId)
	if err != nil {
		log.Printf("user stratus error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "user_status_error",
			"message": fmt.Sprintf("user stratus error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// imageの削除
	if activity.ImageUrl != "" {
		imageFileName := strings.Replace(activity.ImageUrl, config.Config.GcsObjectPath, "", -1)
		fmt.Printf("image file name: %s", imageFileName)
		err = ImageDelete(imageFileName)
		if err != nil {
			log.Printf("image delete error: %v", err)
			return c.JSON(fiber.Map{
				"status":  false,
				"code":    "image_delete_error",
				"message": fmt.Sprintf("image delete error: %v", err),
				"data":    fiber.Map{},
			})
		}
	}

	// activity情報の削除
	err = db.DB.Delete(activity).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_activity_delete",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to delete activity")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_activity_delete",
		"message": "",
		"data":    fiber.Map{},
	})
}
