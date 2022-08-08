package service

import (
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
)

/*
 * GetStatusFromId
 * idからstatus詳細情報を取得
 * @return models.Status
 */
func GetStatusFromId(id string) (models.Status, error) {
	var status models.Status
	err := db.DB.Where("id", id).First(&status).Error
	if err != nil {
		return status, err
	}
	return status, nil
}
