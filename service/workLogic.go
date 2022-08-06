package service

import (
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
)

/*
 * GetWorkFromId
 * idからwork詳細情報を取得
 * @return models.Work
 */
func GetWorkFromId(id string) (models.Work, error) {
	var work models.Work
	err := db.DB.Where("id", id).First(&work).Error
	if err != nil {
		return work, err
	}
	return work, nil
}
