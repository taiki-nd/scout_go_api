package service

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
)

/*
 * GetUserFromId
 * idからuser詳細情報を取得
 * @return models.User
 */
func GetUserFromId(c *fiber.Ctx) (models.User, error) {
	id, _ := strconv.Atoi(c.Params("id"))
	var user models.User
	err := db.DB.Where("id", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
