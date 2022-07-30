package service

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
)

/*
 * GetPrefectureFromId
 * idからprefecture詳細情報を取得
 * @return models.Prefecture
 */
func GetPrefectureFromId(c *fiber.Ctx) (models.Prefecture, error) {
	id, _ := strconv.Atoi(c.Params("id"))
	var prefecture models.Prefecture
	err := db.DB.Where("id", id).First(&prefecture).Error
	if err != nil {
		return prefecture, err
	}
	return prefecture, nil
}
