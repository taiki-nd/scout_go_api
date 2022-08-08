package service

import (
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
)

/*
 * GetProjectFromId
 * idからproject詳細情報を取得
 * @return models.Project
 */
func GetProjectFromId(id string) (models.Project, error) {
	var project models.Project
	err := db.DB.Where("id", id).First(&project).Error
	if err != nil {
		return project, err
	}
	return project, nil
}
