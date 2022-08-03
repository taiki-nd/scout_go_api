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
 * LicensesIndex
 * license一覧を取得
 */
func LicensesIndex(c *fiber.Ctx) error {
	log.Println("start to get licenses")

	var licenses []*models.License

	// licenseesレコードの取得
	err := db.DB.Find(&licenses).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_get_licensees",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    licenses,
		})
	}

	log.Println("success to get licenses")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_license_index",
		"message": "",
		"data":    licenses,
	})
}

/*
 * LicensesCreate
 * DBにlicense情報を登録
 */
func LicensesCreate(c *fiber.Ctx) error {
	log.Println("start to create license")

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

	var license models.License

	// リクエストボディのパース
	err := c.BodyParser(&license)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_parse_license_create",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// レコード作成
	err = db.DB.Create(&license).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_license_create",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to create license")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_license_create",
		"message": "",
		"data":    license,
	})
}

/*
 * LicensesShow
 * license詳細情報の取得
 */
func LicensesShow(c *fiber.Ctx) error {
	log.Println("start to get license")

	id := c.Params("id")

	// license情報の取得
	license, err := service.GetLicenseFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_license_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Printf("success to get licenses")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_license_show",
		"message": "",
		"data":    license,
	})
}

/*
 * LicensesUpdate
 * license情報の更新
 */
func LicensesUpdate(c *fiber.Ctx) error {
	log.Println("start to Update license")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// license情報の取得
	license, err := service.GetLicenseFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_license_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// userステータスの確認
	err = service.CheckUserStatus(uuid, license.UserId)
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
	err = c.BodyParser(&license)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_parse_license_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// license情報の更新
	err = db.DB.Model(&license).Updates(license).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_license_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to update license")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_license_update",
		"message": fmt.Sprintf("db error: %v", err),
		"data":    license,
	})
}

/*
 * LicensesDelete
 * license情報の削除
 */
func LicensesDelete(c *fiber.Ctx) error {
	log.Println("start to delete license")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// license情報の取得
	license, err := service.GetLicenseFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_license_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// userステータスの確認
	err = service.CheckUserStatus(uuid, license.UserId)
	if err != nil {
		log.Printf("user stratus error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "user_status_error",
			"message": fmt.Sprintf("user stratus error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// license情報の削除
	err = db.DB.Delete(license).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_license_delete",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to delete license")

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_license_delete",
		"message": "",
		"data":    fiber.Map{},
	})
}
