package service

import (
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
)

/*
 * GetResumeFromId
 * idからresume詳細情報を取得
 * @return models.Resume
 */
func GetResumeFromId(id string) (models.Resume, error) {
	var resume models.Resume
	err := db.DB.Where("id", id).First(&resume).Error
	if err != nil {
		return resume, err
	}
	return resume, nil
}
