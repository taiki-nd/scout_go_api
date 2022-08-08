package service

import (
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
)

/*
 * GetActivityFromId
 * idからactivity詳細情報を取得
 * @return models.Activity
 */
func GetActivityFromId(id string) (models.Activity, error) {
	var activity models.Activity
	err := db.DB.Where("id", id).First(&activity).Error
	if err != nil {
		return activity, err
	}
	return activity, nil
}
