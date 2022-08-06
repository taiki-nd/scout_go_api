package service

import (
	"fmt"
	"log"

	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
)

/*
 * GetUserFromId
 * idからuser詳細情報を取得
 * @return models.User
 */
func GetUserFromId(id string) (models.User, error) {
	var user models.User
	err := db.DB.Preload("Statuses").Preload("Prefectures").Preload("Schools").Preload("Activities").Where("id", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

/*
 * GetUserFromUuid
 * idからuser詳細情報を取得
 * @return models.User
 */
func GetUserFromUuid(uuid string) (models.User, error) {
	var user models.User
	err := db.DB.Where("uuid", uuid).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

/*
 * CheckUserStatus
 * userの管理情報を精査
 * @params user models.User
 */
func CheckUserStatus(uuid string, userId uint) error {
	// サインイン状態の確認
	if uuid == "" {
		log.Println("not signin")
		return fmt.Errorf("not_signin")
	}

	// admin権限の確認
	signinUser, err := GetUserFromUuid(uuid)
	if err != nil {
		log.Println("db error: GetUserFromUuid()")
		return fmt.Errorf("db error/ GetUserFromUuid()")
	}
	if signinUser.IsAdmin {
		return nil
	}

	// サインインユーザーの一致確認
	if userId == signinUser.Id {
		return nil
	} else {
		return fmt.Errorf("not match user")
	}
}

/*
 * CheckAdmin
 * userのadmin情報を確認
 */
func CheckAdmin(uuid string) error {
	// サインイン状態の確認
	if uuid == "" {
		log.Println("not signin")
		return fmt.Errorf("not_signin")
	}

	// admin権限の確認
	signinUser, err := GetUserFromUuid(uuid)
	if err != nil {
		log.Println("db error: GetUserFromUuid()")
		return fmt.Errorf("db error/ GetUserFromUuid()")
	}
	if signinUser.IsAdmin {
		return nil
	} else {
		return fmt.Errorf("user is not admin")
	}
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
