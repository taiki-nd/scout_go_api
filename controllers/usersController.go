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
 * UsersIndex
 * user一覧を取得
 */
func UsersIndex(c *fiber.Ctx) error {
	log.Println("start to get users")

	var users []*models.User

	// usersレコードの取得
	err := db.DB.Preload("Statuses").Preload("Prefectures").Preload("Schools").Preload("Activities").Preload("Works").Find(&users).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_get_users",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    users,
		})
	}

	log.Println("success to get users")

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

	var userAssociation models.UserAssociation
	log.Println(userAssociation)

	// リクエストボディのパース
	err := c.BodyParser(&userAssociation)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_parse_user_create",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	statuses := service.GetStatuses(userAssociation.Statuses)
	prefectures := service.GetPrefectures(userAssociation.Prefectures)

	user := models.User{
		Uuid:           userAssociation.Uuid,
		LastName:       userAssociation.LastName,
		LastNameKana:   userAssociation.LastNameKana,
		FirstName:      userAssociation.FirstName,
		FirstNameKana:  userAssociation.FirstNameKana,
		Nickname:       userAssociation.Nickname,
		Sex:            userAssociation.Sex,
		BirthYear:      userAssociation.BirthYear,
		BirthMonth:     userAssociation.BirthMonth,
		BirthDay:       userAssociation.BirthDay,
		AutoPermission: userAssociation.AutoPermission,
		IsExample:      userAssociation.IsExample,
		IsAdmin:        userAssociation.IsAdmin,
		Statuses:       statuses,
		Prefectures:    prefectures,
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

	id := c.Params("id")

	// user情報の取得
	user, err := service.GetUserFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_user_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Printf("success to get users: %v", user.Nickname)

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

	id := c.Params("id")
	uuid := c.Query("uuid")

	// user情報の取得
	user, err := service.GetUserFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_user_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// userステータスの確認
	err = service.CheckUserStatus(uuid, user.Id)
	if err != nil {
		log.Printf("user stratus error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "user_status_error",
			"message": fmt.Sprintf("user stratus error: %v", err),
			"data":    fiber.Map{},
		})
	}

	var userAssociation models.UserAssociation

	// リクエストボディのパース
	err = c.BodyParser(&userAssociation)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_parse_user_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// transaction開始
	errTransaction := db.DB.Transaction(func(tx *gorm.DB) error {
		// アソシエーションの削除
		errStatus := tx.Table("user_statuses").Where("user_id = ?", user.Id).Delete("").Error
		if errStatus != nil {
			log.Printf("db error: %v", errStatus)
			return fmt.Errorf("db error: %v", errStatus)
		}
		errPrefecture := tx.Table("user_prefectures").Where("user_id = ?", user.Id).Delete("").Error
		if errPrefecture != nil {
			log.Printf("db error: %v", errPrefecture)
			return fmt.Errorf("db error: %v", errPrefecture)
		}

		// アソシエーションの更新
		statuses := service.GetStatuses(userAssociation.Statuses)
		prefectures := service.GetPrefectures(userAssociation.Prefectures)

		userForUpdate := models.User{
			Id:             user.Id,
			Uuid:           userAssociation.Uuid,
			LastName:       userAssociation.LastName,
			LastNameKana:   userAssociation.LastNameKana,
			FirstName:      userAssociation.FirstName,
			FirstNameKana:  userAssociation.FirstNameKana,
			Nickname:       userAssociation.Nickname,
			Sex:            userAssociation.Sex,
			BirthYear:      userAssociation.BirthYear,
			BirthMonth:     userAssociation.BirthMonth,
			BirthDay:       userAssociation.BirthDay,
			AutoPermission: userAssociation.AutoPermission,
			IsExample:      userAssociation.IsExample,
			IsAdmin:        userAssociation.IsAdmin,
			Statuses:       statuses,
			Prefectures:    prefectures,
		}

		// user情報の更新
		err = tx.Model(&userForUpdate).Updates(userForUpdate).Error
		if err != nil {
			log.Printf("db error: %v", err)
			return fmt.Errorf("db error: %v", errStatus)
		}
		return nil
	})
	if errTransaction != nil {
		log.Println(errTransaction)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_user_update",
			"message": fmt.Sprintf("db error: %v", errTransaction),
			"data":    fiber.Map{},
		})
	}

	log.Printf("success to update user: %v", user.Nickname)

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_user_update",
		"message": "",
		"data":    fiber.Map{},
	})
}

/*
 * UserDelete
 * user情報の削除
 */
func UsersDelete(c *fiber.Ctx) error {
	log.Println("start to delete user")

	id := c.Params("id")
	uuid := c.Query("uuid")

	log.Printf("uuid: %v", uuid)

	// user情報の取得
	user, err := service.GetUserFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_user_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// userステータスの確認
	err = service.CheckUserStatus(uuid, user.Id)
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
		// アソシエーションの削除
		errStatus := tx.Table("user_statuses").Where("user_id = ?", user.Id).Delete("").Error
		if errStatus != nil {
			log.Printf("db error: %v", errStatus)
			return fmt.Errorf("db error: %v", errStatus)
		}
		errPrefecture := tx.Table("user_prefectures").Where("user_id = ?", user.Id).Delete("").Error
		if errPrefecture != nil {
			log.Printf("db error: %v", errPrefecture)
			return fmt.Errorf("db error: %v", errPrefecture)
		}
		errLicense := tx.Table("licenses").Where("user_id = ?", user.Id).Delete("").Error
		if errLicense != nil {
			log.Printf("db error: %v", errLicense)
			return fmt.Errorf("db error: %v", errLicense)
		}
		errSchool := tx.Table("schools").Where("user_id = ?", user.Id).Delete("").Error
		if errSchool != nil {
			log.Printf("db error: %v", errSchool)
			return fmt.Errorf("db error: %v", errSchool)
		}
		errActivity := tx.Table("activities").Where("user_id = ?", user.Id).Delete("").Error
		if errActivity != nil {
			log.Printf("db error: %v", errActivity)
			return fmt.Errorf("db error: %v", errActivity)
		}
		var work_ids []int64
		tx.Table("works").Where("user_id = ?", user.Id).Pluck("id", &work_ids)
		if len(work_ids) != 0 {
			err := tx.Table("projects").Where("work_id IN (?)", work_ids).Delete("").Error
			if err != nil {
				return fmt.Errorf("db error: %v", err)
			}
		}
		errWork := tx.Table("works").Where("user_id = ?", user.Id).Delete("").Error
		if errWork != nil {
			log.Printf("db error: %v", errWork)
			return fmt.Errorf("db error: %v", errWork)
		}
		errResume := tx.Table("resumes").Where("user_id = ?", user.Id).Delete("").Error
		if errResume != nil {
			log.Printf("db error: %v", errResume)
			return fmt.Errorf("db error: %v", errResume)
		}

		// user情報の削除
		err = tx.Delete(user).Error
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
			"code":    "failed_db_user_delete",
			"message": fmt.Sprintf("db error: %v", errTransaction),
			"data":    fiber.Map{},
		})
	}

	log.Printf("success to delete user: %v", user.Nickname)

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_user_delete",
		"message": "",
		"data":    fiber.Map{},
	})
}

/*
 * UserFromUuid
 * Uuidからuser情報を取得
 */
func UserFromUuid(c *fiber.Ctx) error {
	uuid := c.Query("uuid")

	// user情報の取得
	_, err := service.GetUserFromUuid(uuid)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_get_user_from_uuid",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"code":    "success_get_user_from_uuid",
		"message": "",
		"data":    fiber.Map{},
	})
}
