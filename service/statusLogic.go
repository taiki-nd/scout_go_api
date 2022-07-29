package service

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
)

/*
 * GetStatusFromId
 * idからstatus詳細情報を取得
 * @return models.Status
 */
func GetStatusFromId(c *fiber.Ctx) (models.Status, error) {
	id, _ := strconv.Atoi(c.Params("id"))
	var status models.Status
	err := db.DB.Where("id", id).First(&status).Error
	if err != nil {
		return status, err
	}
	return status, nil
}
