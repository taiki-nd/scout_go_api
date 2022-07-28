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
	err := db.DB.Preload("Statuses").Find(&users).Error
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

	log.Printf("user association:	%v", userAssociation.LastName)

	statuses := service.GetStatuses(userAssociation.Statuses)

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

		// アソシエーションの更新
		statuses := service.GetStatuses(userAssociation.Statuses)

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

	// user情報の削除
	err = db.DB.Delete(user).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "failed_db_user_delete",
			"message": fmt.Sprintf("db error: %v", err),
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
