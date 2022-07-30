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
	err := db.DB.Preload("Statuses").Preload("Prefectures").Where("id", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

/*
 * GetStatuses
 * idからstatus情報を取得
 * @params statusesIds
 * @return []models.Status
 */
func GetStatuses(statusesIds []int) []models.Status {
	statuses := make([]models.Status, len(statusesIds))
	for i, statusId := range statusesIds {
		statuses[i] = models.Status{
			Id: uint(statusId),
		}
	}
	return statuses
}

/*
 * Get prefectures
 * idからprefecture情報を取得
 * @params prefecturesIds
 * @return []models.Prefecture
 */
func GetPrefectures(prefecturesIds []int) []models.Prefecture {
	prefectures := make([]models.Prefecture, len(prefecturesIds))
	for i, prefectureId := range prefecturesIds {
		prefectures[i] = models.Prefecture{
			Id: uint(prefectureId),
		}
	}
	return prefectures
}
