package service

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
)

/*
 * GetLicenseFromId
 * idからlicense詳細情報を取得
 * @return models.License
 */
func GetLicenseFromId(c *fiber.Ctx) (models.License, error) {
	id, _ := strconv.Atoi(c.Params("id"))
	var license models.License
	err := db.DB.Where("id", id).First(&license).Error
	if err != nil {
		return license, err
	}
	return license, nil
}
