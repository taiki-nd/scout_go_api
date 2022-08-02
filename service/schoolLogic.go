package service

import (
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
)

/*
 * GetSchoolFromId
 * idからschool詳細情報を取得
 * @return models.School
 */
func GetSchoolFromId(id string) (models.School, error) {
	var school models.School
	err := db.DB.Where("id", id).First(&school).Error
	if err != nil {
		return school, err
	}
	return school, nil
}
